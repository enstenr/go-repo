package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	chatv1 "github.com/enstenr/go-repo/gen/pb/chat/v1"
	"github.com/enstenr/go-repo/gen/pb/chat/v1/chatv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ChatServer struct {
	userClient chatv1connect.UserServiceClient

	// A map to keep track of active streams (Key: UserID)
	// In a real app, this would be Redis, but for a showcase, a Map works!
	mu          sync.RWMutex
	bidiClients map[string]*connect.BidiStream[chatv1.Message, chatv1.Message]
	clients     map[string]chan *chatv1.Message
}

func (s *ChatServer) Connect(ctx context.Context,
	stream *connect.BidiStream[chatv1.Message, chatv1.Message],
) error {

	log.Println(" Connection established...)")
	for {
		msg, err := stream.Receive()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(" %s %s ", msg.SenderId, msg.Text)
		s.mu.Lock()
		s.bidiClients[msg.SenderId] = stream
		s.mu.Unlock()

		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}
func (s *ChatServer) Subscribe(
	ctx context.Context,
	req *connect.Request[chatv1.SubscribeRequest],
	stream *connect.ServerStream[chatv1.Message],
) error {
	userID := req.Msg.UserId
	msgChan := make(chan *chatv1.Message, 10)

	s.mu.Lock()
	s.clients[userID] = msgChan
	s.mu.Unlock()

	log.Printf("User %s subscribed for messages", userID)

	defer func() {
		s.mu.Lock()
		delete(s.clients, userID)
		s.mu.Unlock()
		close(msgChan)
		log.Printf("User %s disconnected", userID)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-msgChan:
			if err := stream.Send(msg); err != nil {
				return err
			}
		}
	}
}

// 2. SEND MESSAGE (For React - Unary)
func (s *ChatServer) SendMessage(
	ctx context.Context,
	req *connect.Request[chatv1.Message],
) (*connect.Response[chatv1.SendMessageResponse], error) {
	msg := req.Msg
	log.Printf("Incoming message from %s to %s: %s", msg.SenderId, msg.RecipientId, msg.Text)
	_, err := s.userClient.GetUser(ctx, connect.NewRequest(&chatv1.GetUserRequest{
		UserId: msg.RecipientId,
	}))
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	// If no recipient specified, echo back to sender for testing
	if recipientChan, ok := s.clients[msg.RecipientId]; ok {
		recipientChan <- msg
	} else {
		// Return an actual error to React
		return nil, connect.NewError(connect.CodeNotFound, errors.New("User is offline"))
	}
	if senderChan, ok := s.clients[msg.SenderId]; ok && msg.SenderId != msg.RecipientId {
		senderChan <- msg
	}

	return connect.NewResponse(&chatv1.SendMessageResponse{MsgId: "delivered"}), nil
}

func main() {

	userSvcClient := chatv1connect.NewUserServiceClient(
		http.DefaultClient,
		"http://localhost:50051",
	)

	chatServer := &ChatServer{
		userClient:  userSvcClient,
		clients:     make(map[string]chan *chatv1.Message),
		bidiClients: make(map[string]*connect.BidiStream[chatv1.Message, chatv1.Message]),
	}
	mux := http.NewServeMux()
	path, handler := chatv1connect.NewChatServiceHandler(chatServer)
	mux.Handle(path, handler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{
			"Connect-Protocol-Version",
			"Content-Type",
			"Connect-Timeout",
			"X-User-Agent",
			"X-Connect-Protocol-Version",
		},
		ExposedHeaders: []string{
			"Connect-Error-Message",
			"Connect-Status-Details",
		},
		Debug: true, // Turn this off in production
	})
	// 3. Wrap the handler
	// We wrap h2c (which handles HTTP/2) with the CORS middleware
	handlerWithCors := c.Handler(h2c.NewHandler(mux, &http2.Server{}))
	log.Println(" Server starting... 50052")
	err := http.ListenAndServe("localhost:50052", handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}

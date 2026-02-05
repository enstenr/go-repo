package main

import (
	"context"
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

	s.mu.RLock()
	// If no recipient specified, echo back to sender for testing
	target := msg.RecipientId
	if target == "" {
		target = msg.SenderId
	}

	if ch, ok := s.clients[target]; ok {
		ch <- msg
	}
	s.mu.RUnlock()

	return connect.NewResponse(&chatv1.SendMessageResponse{MsgId: "delivered"}), nil
}

func main() {
	chatServer := &ChatServer{
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

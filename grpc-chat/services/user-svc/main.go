package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	chatv1 "github.com/enstenr/go-repo/gen/pb/chat/v1"
	"github.com/enstenr/go-repo/gen/pb/chat/v1/chatv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UserServer struct{}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *connect.Request[chatv1.GetUserRequest],
) (*connect.Response[chatv1.User], error) {

	log.Printf("Incoming request for User ID: %s", req.Msg.UserId)

	// In a real app, you would fetch from a database here.
	// We'll return a hardcoded user for the showcase.
	user := &chatv1.User{
		Id:     req.Msg.UserId,
		Name:   "Architect Rajesh",
		Status: "Available",
	}

	// Wrap the protobuf message in a Connect response
	return connect.NewResponse(user), nil
}
func main() {
	userServer := &UserServer{}
	mux := http.NewServeMux()

	// 3. Register the service handler on the mux
	// This creates the URL path: /chat.v1.UserService/GetUser
	path, handler := chatv1connect.NewUserServiceHandler(userServer)
	mux.Handle(path, handler)

	log.Println("User Service is starting on localhost:50051...")

	// 4. Start the server
	// h2c.NewHandler enables HTTP/2 without requiring TLS certificates for local dev
	err := http.ListenAndServe(
		"localhost:50051",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

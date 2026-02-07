package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	chatv1 "github.com/enstenr/go-repo/gen/pb/chat/v1"
	"github.com/enstenr/go-repo/gen/pb/chat/v1/chatv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UserServer struct {
	users map[string]*chatv1.User
}

func NewUserServer() *UserServer {
	return &UserServer{
		users: make(map[string]*chatv1.User),
	}
}
func (s *UserServer) GetUser(
	ctx context.Context,
	req *connect.Request[chatv1.GetUserRequest],
) (*connect.Response[chatv1.User], error) {

	log.Printf("Incoming request for User ID: %s", req.Msg.UserId)

	user, ok := s.users[req.Msg.UserId]
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, nil)
	}
	return connect.NewResponse(user), nil
}
func (s *UserServer) Register(
	ctx context.Context,
	req *connect.Request[chatv1.RegisterRequest],
) (*connect.Response[chatv1.RegisterResponse], error) {

	// Generate a simple ID (in production, use a UUID)
	newID := "user_" + req.Msg.Name

	newUser := &chatv1.User{
		Id:     newID,
		Name:   req.Msg.Name,
		Status: "Just Joined",
	}

	// Save to our "database"
	s.users[newID] = newUser

	log.Printf("New user registered: %s with ID: %s", newUser.Name, newUser.Id)

	return connect.NewResponse(&chatv1.RegisterResponse{
		User: newUser,
	}), nil
}
func (s *UserServer) ListUsers(
	ctx context.Context,
	req *connect.Request[chatv1.ListUsersRequest],
) (*connect.Response[chatv1.ListUsersResponse], error) {

	log.Println("Fetching all users...")

	// Convert our map values into a slice
	allUsers := make([]*chatv1.User, 0, len(s.users))
	for _, user := range s.users {
		allUsers = append(allUsers, user)
	}

	return connect.NewResponse(&chatv1.ListUsersResponse{
		Users: allUsers,
	}), nil
}

func main() {

	userServer := NewUserServer()
	mux := http.NewServeMux()

	// 3. Register the service handler on the mux
	// This creates the URL path: /chat.v1.UserService/GetUser
	path, handler := chatv1connect.NewUserServiceHandler(userServer)
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
	log.Println("User Service is starting on localhost:50051...")

	/// 3. Wrap the handler
	// We wrap h2c (which handles HTTP/2) with the CORS middleware
	handlerWithCors := c.Handler(h2c.NewHandler(mux, &http2.Server{}))
	log.Println(" Server starting... 50051")
	err := http.ListenAndServe("localhost:50051", handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import ( "context"
"log"
	"net"
"time"


	"google.golang.org/grpc"
"go-grpc-metadata-service/pb"
)
import "google.golang.org/protobuf/proto"
type UserServer struct{
pb.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserProfile, error) {
	// For now, let's just return a hardcoded user
	return &pb.UserProfile{

		Email: "srajesh2712@gmail.com",
		Id:    req.UserId,
        		Name:  "Rajesh",
	}, nil
}
func (s *UserServer) GetUserStream(req *pb.UserRequest, stream pb.UserService_GetUserStreamServer) error {
	log.Printf("Streaming users starting from ID: %d", req.UserId)

	// We'll simulate a "buffer" of multiple users
	users := []*pb.UserProfile{
		{Id: req.UserId, Name: "Rajesh", Email: "rajesh@example.com"},
		{Id: req.UserId + 1, Name: "Alice", Email: "alice@example.com"},
		{Id: req.UserId + 2, Name: "Bob", Email: "bob@example.com"},
	}

	for _, user := range users {
		// Simulating a delay/processing time for each "chunk"
		time.Sleep(1 * time.Second)

		// This sends one "frame" over the wire immediately
		if err := stream.Send(user); err != nil {
			return err
		}
		log.Printf("Sent chunk for User: %s", user.Name)
	}

	// Returning nil tells gRPC "The stream is finished"
	return nil
}

func (s *UserServer) GetUserBinary(req *pb.UserRequest, stream pb.UserService_GetUserBinaryServer) error {
	// 1. Form the local array
    	userList := &pb.UserList{} // Ensure you have this message in proto
    	for i := 0; i < 1000; i++ {
    		userList.Users = append(userList.Users, &pb.UserProfile{Id: int32(i), Name: "Bulk"})
    	}

    	// 2. Serialize the WHOLE array to bytes
    	fullBinary, _ := proto.Marshal(userList)

    	// 3. Define Chunk Size (e.g., 64KB)
    	chunkSize := 64 * 1024
    	totalSize := len(fullBinary)

    	for i := 0; i < totalSize; i += chunkSize {
    		end := i + chunkSize
    		if end > totalSize {
    			end = totalSize
    		}

    		// 4. Stream the slice
    		err := stream.Send(&pb.BinaryPayload{Data: fullBinary[i:end]})
    		if err != nil {
    			return err
    		}
    	}
    	return nil
}

func main() {
lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 3. THE DISPATCHER (The gRPC Engine)
	grpcServer := grpc.NewServer()

	// Connect the Logic to the Engine
	pb.RegisterUserServiceServer(grpcServer, &UserServer{})

	log.Println("Server is running on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
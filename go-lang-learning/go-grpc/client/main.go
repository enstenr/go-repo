package main

import (
	"context"
	"log"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go-grpc-metadata-service/pb" // Our contract
)
import "google.golang.org/protobuf/proto"
func main() {
	// 1. Dial the server (Establish the physical "pipe")
	// We use insecure credentials for this local hands-on
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. Create the Stub (The EJB "Remote Reference" equivalent)
	client := pb.NewUserServiceClient(conn)

	// 3. Prepare a request and a timeout (Context)
	// As an architect, you'll love Context - it handles distributed timeouts!
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 4. Make the call
	resp, err := client.GetUser(ctx, &pb.UserRequest{UserId: 101})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("Response from Server: Name=%s, Email=%s", resp.Name, resp.Email)


	// --- CALL 2: STREAMING ---
    	log.Println("\n--- Calling Streaming GetUserStream ---")
    	stream, _ := client.GetUserStream(context.Background(), &pb.UserRequest{UserId: 200})
    	for {
    		user, err := stream.Recv()
    		if err == io.EOF {
    			break
    		}
    		log.Printf("Stream Buffer Received: %s (ID: %d)", user.Name, user.Id)
    	}


  // 1. Call the binary method.
  // Use a fresh variable name 'binStream' so the compiler knows the type is BinaryPayload
  binStream, err := client.GetUserBinary(context.Background(), &pb.UserRequest{})
  if err != nil {
      log.Fatalf("Error: %v", err)
  }
var receivedBlob []byte // The "reconstruction area"

	for {
		chunk, err := binStream.Recv()
		if err == io.EOF {
			break // All pieces arrived
		}
		// APPEND: Building the binary blob piece by piece
		receivedBlob = append(receivedBlob, chunk.Data...)
	}

	// 5. RECONSTRUCT: Convert the finished blob back into the Array Object
	finalList := &pb.UserList{}
	if err := proto.Unmarshal(receivedBlob, finalList); err != nil {
		log.Fatal("Failed to reconstruct array:", err)
	}
log.Printf("Successfully reconstructed array with %d users", len(finalList.Users))
}


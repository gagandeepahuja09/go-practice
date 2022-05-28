package main

import (
	"context"
	"log"

	pb "dist-store.com/keyvalue"
)

type server struct {
	pb.UnimplementedKeyValueServer
}

// note: here we need not worry about writing the response to the response writer.
func (s *server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received GET key=%v", r.Key)
	value, err := Get(r.Key)
	return &pb.GetResponse{Value: value}, err
}

// to run this: comment the other implements of main methods in the same directory
// func main() {
// 	s := grpc.NewServer()
// 	pb.RegisterKeyValueServer(s, &server{})

// 	// Open a listening port on 50051
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	// Start accepting connections on the listening port
// 	if err = s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

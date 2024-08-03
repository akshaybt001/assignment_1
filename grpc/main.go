package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "workout3/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMockServiceServer
}

func (s *server) GetSomeData(ctx context.Context, req *pb.UserData) (*pb.UserData, error) {
	res := fmt.Sprintf("welcome %s", req.Name)
	return &pb.UserData{Name: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMockServiceServer(s, &server{})
	log.Println("gRPC listening on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

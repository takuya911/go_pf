package main

import (
	"context"
	"log"
	"net"

	pb "github.com/takuya911/go_pf/services/user/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// GET User
func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("request received")
	var id = in.Id
	var name = "name"
	return &pb.GetUserResponse{Id: id, Name: name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

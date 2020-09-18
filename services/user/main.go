package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/takuya911/go_pf/services/user/infrastructure"
	pb "github.com/takuya911/go_pf/services/user/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// GET User
func (s *server) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	log.Printf("request received")
	ts := ptypes.TimestampNow()
	result := &pb.User{
		Id:        in.Id,
		Name:      "name",
		Email:     "email@email.com",
		Password:  "password",
		CreatedAt: ts,
		UpdatedAt: ts,
		DeletedAt: ts,
	}
	return result, nil
}

func main() {
	// DB Connect
	dbConn, err := infrastructure.NewGormConnect()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

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

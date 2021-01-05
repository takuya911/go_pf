package infrastructure

import (
	pb "github.com/takuya911/gqlgen-grpc/services/user/proto"
	"google.golang.org/grpc"
)

// NewGrpcServer function
func NewGrpcServer(uc pb.UserServiceServer) *grpc.Server {
	server := grpc.NewServer()
	return server
}

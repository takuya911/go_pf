package infrastructure

import (
	pb "github.com/takuya911/go_pf/services/user/proto"
	"google.golang.org/grpc"
)

// NewGrpcServer function
func NewGrpcServer(uc pb.UserServiceServer) *grpc.Server {
	server := grpc.NewServer()
	return server
}

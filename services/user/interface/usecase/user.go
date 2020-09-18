package usecase

import (
	"context"

	pb "github.com/takuya911/go_pf/services/user/proto"
)

// UserUsecase interface
type UserUsecase interface {
	GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error)
}

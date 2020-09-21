package usecase

import (
	"context"

	pb "github.com/takuya911/go_pf/services/user/proto"

	"github.com/takuya911/go_pf/services/user/domain"
)

// UserUsecase interface
type UserUsecase interface {
	GetUser(ctx context.Context, in *pb.GetUserReq) (*domain.User, error)
}

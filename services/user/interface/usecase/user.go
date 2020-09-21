package usecase

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

// UserUsecase interface
type UserUsecase interface {
	GetUserByID(ctx context.Context, in *pb.GetUserReq) (*domain.User, error)
}

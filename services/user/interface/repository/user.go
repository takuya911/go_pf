package repository

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

// UserRepository interface
type UserRepository interface {
	GetUser(ctx context.Context, in *pb.GetUserReq) (*domain.User, error)
}

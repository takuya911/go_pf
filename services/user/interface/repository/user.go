package repository

import (
	"context"

	pb "github.com/takuya911/go_pf/services/user/proto"
)

// UserRepository interface
type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*pb.User, error)
}

package repository

import (
	"context"

	"github.com/takuya911/gqlgen-grpc/services/user/domain"
)

// UserRepository interface
type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int64) (bool, error)
}

package usecase

import (
	"context"

	"github.com/takuya911/gqlgen-grpc/services/user/domain"
)

// UserUsecase interface
type UserUsecase interface {
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
	Login(ctx context.Context, email string, password string) (*domain.User, *domain.TokenPair, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.TokenPair, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, *domain.User, error)
	DeleteUser(ctx context.Context, user *domain.User) (bool, error)
}

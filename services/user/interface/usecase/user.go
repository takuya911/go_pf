package usecase

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
)

// UserUsecase interface
type UserUsecase interface {
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
}

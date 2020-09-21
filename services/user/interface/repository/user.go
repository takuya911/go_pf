package repository

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
)

// UserRepository interface
type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
}

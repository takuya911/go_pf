package interactor

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
	"github.com/takuya911/go_pf/services/user/usecase/repository"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

// NewUserInteractor function
func NewUserInteractor(u repository.UserRepository) *userInteractor {
	return &userInteractor{u}
}

// UserUsecase interface
type UserUsecase interface {
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
}

func (i *userInteractor) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	result, err := i.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

package interactor

import (
	"context"
	"log"
	"strconv"

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

func (i *userInteractor) Login(ctx context.Context, email string, password string) (*domain.User, *domain.TokenPair, error) {
	user, err := i.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	if err := compareHashAndPass(user.Password, password); err != nil {
		log.Println(err)
		return nil, nil, err
	}

	tokenPair, err := genTokenPair(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}

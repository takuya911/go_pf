package usecase

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
	"github.com/takuya911/go_pf/services/user/interface/repository"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

// NewUserInteractor function
func NewUserInteractor(u repository.UserRepository) *userInteractor {
	return &userInteractor{u}
}

func (i *userInteractor) GetUser(ctx context.Context, in *pb.GetUserReq) (*domain.User, error) {
	result, err := i.userRepository.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

// get user
func (i *userInteractor) GetUser(ctx context.Context, in *pb.GetUserReq) (*domain.User, error) {
	result := &domain.User{}
	return result, nil
}

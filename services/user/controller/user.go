package controller

import (
	"context"
	"time"

	"github.com/takuya911/go_pf/services/user/domain"
	"github.com/takuya911/go_pf/services/user/interface/usecase"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

type userController struct {
	userInteractor usecase.UserUsecase
}

// NewUserController function
func NewUserController(u usecase.UserUsecase) pb.UserServiceServer {
	return &userController{u}
}

func (c *userController) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	result := &domain.User{
		ID:        in.Id,
		Name:      "name",
		Email:     "email@email.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}
	return convUserProto(result)
}

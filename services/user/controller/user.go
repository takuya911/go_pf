package controller

import (
	"context"
	"log"

	"github.com/takuya911/go_pf/services/user/errors"
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

func (c *userController) GetUserByID(ctx context.Context, in *pb.GetUserForm) (*pb.User, error) {
	userID := in.GetId()
	if userID == 0 {
		return nil, errors.BadRequestError
	}

	result, err := c.userInteractor.GetUserByID(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	if result.ID == 0 {
		return &pb.User{}, errors.BadRequestError
	}
	return convUserProto(result)
}

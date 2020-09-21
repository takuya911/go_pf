package controller

import (
	"context"
	"log"

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
	result, err := c.userInteractor.GetUser(ctx, in)
	if err != nil {
		log.Fatal(err)
	}
	return convUserProto(result)
}

package controller

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"
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
	log.Printf("request received")
	ts := ptypes.TimestampNow()
	result := &pb.User{
		Id:        in.Id,
		Name:      "name",
		Email:     "email@email.com",
		Password:  "password",
		CreatedAt: ts,
		UpdatedAt: ts,
		DeletedAt: ts,
	}
	return result, nil
}

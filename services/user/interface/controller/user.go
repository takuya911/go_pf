package controller

import (
	"context"

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
		return nil, err
	}
	return convUserProto(result)
}

func (c *userController) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	user, token, err := c.userInteractor.Login(ctx, in.Email, in.Password)

	if err != nil {
		return nil, err
	}
	userproto, err := convUserProto(user)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{User: userproto, TokenPair: convTokenPairProto(token)}, nil
}

func (c *userController) CreateUser(stream pb.UserService_CreateUserServer) error {
	return stream.SendAndClose(&pb.CreateUserRes{
		User:      &pb.User{},
		TokenPair: &pb.TokenPair{},
	})
}

package controller

import (
	"context"

	"github.com/takuya911/go_pf/services/user/domain"
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
		return nil, errors.PasswordFaultError
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
	ctx := stream.Context()

	req, err := stream.Recv()
	if err != nil {
		return err
	}
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := c.userInteractor.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	pbUser, err := convUserProto(user)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.CreateUserRes{
		User:      pbUser,
		TokenPair: convTokenPairProto(token),
	})
}

func (c *userController) UpdateUser(stream pb.UserService_UpdateUserServer) error {
	// token使って認証したい(フロント作ったら)
	ctx := stream.Context()

	req, err := stream.Recv()
	if err != nil {
		return err
	}

	// 更新前ユーザーの取得
	beforeUser, err := c.userInteractor.GetUserByID(ctx, req.Id)
	if err != nil {
		return err
	}

	formUser := &domain.User{
		ID:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// 更新
	if err = c.userInteractor.UpdateUser(ctx, formUser); err != nil {
		return err
	}

	// 更新前ユーザーの取得
	afterUser, err := c.userInteractor.GetUserByID(ctx, req.Id)
	if err != nil {
		return err
	}

	// 変換
	beforeUserProto, err := convUserProto(beforeUser)
	if err != nil {
		return err
	}
	afterUserProto, err := convUserProto(afterUser)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.UpdateUserRes{
		BeforeUser: beforeUserProto,
		AfterUser:  afterUserProto,
	})
}

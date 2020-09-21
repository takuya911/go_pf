package controller

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/takuya911/go_pf/services/user/domain"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

func convUserProto(u *domain.User) (*pb.User, error) {
	createdAt, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	deletedAt, err := ptypes.TimestampProto(u.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil
}

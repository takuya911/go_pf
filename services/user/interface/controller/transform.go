package controller

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/takuya911/gqlgen-grpc/services/user/domain"
	pb "github.com/takuya911/gqlgen-grpc/services/user/proto"
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

	// deleted_atも返そうと思えば返せる?
	// var deletedAt *timestamppb.Timestamp
	// if u.DeletedAt != nil {
	// 	deletedAt, err = ptypes.TimestampProto(*u.DeletedAt)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &pb.User{
		Id:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		Password:        u.Password,
		TelephoneNumber: u.TelephoneNumber,
		Gender:          u.Gender,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		// DeletedAt:       deletedAt,
	}, nil
}

func convTokenPairProto(t *domain.TokenPair) *pb.TokenPair {
	return &pb.TokenPair{
		IdToken: t.IDToken,
	}
}

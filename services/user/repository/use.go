package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

type userRepository struct {
	conn *gorm.DB
}

// NewUserRepository function
func NewUserRepository(conn *gorm.DB) *userRepository {
	return &userRepository{conn}
}

func (r *userRepository) GetUser(ctx context.Context, id int64) (*pb.User, error) {
	var u pb.User
	return &u, nil
}

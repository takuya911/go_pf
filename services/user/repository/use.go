package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/takuya911/go_pf/services/user/domain"
	pb "github.com/takuya911/go_pf/services/user/proto"
)

type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository function
func NewUserRepository(conn *gorm.DB) *userRepository {
	return &userRepository{conn}
}

func (r *userRepository) GetUser(ctx context.Context, in *pb.GetUserReq) (*domain.User, error) {
	var user domain.User
	if result := r.Conn.Where("id = ?", in.GetId()).Find(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

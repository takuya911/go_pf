package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/takuya911/go_pf/services/user/domain"
)

type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository function
func NewUserRepository(conn *gorm.DB) *userRepository {
	return &userRepository{conn}
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	var user domain.User
	if result := r.Conn.Where("id = ?", userID).Find(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

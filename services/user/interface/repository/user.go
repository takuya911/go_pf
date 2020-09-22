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
	if rs := r.Conn.First(&user, userID); rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if rs := r.Conn.Where("email = ?", email).First(&user); rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

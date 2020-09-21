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
	r.Conn.First(&user, userID)
	// If you return errors here, the connection will be lost.
	return &user, nil
}

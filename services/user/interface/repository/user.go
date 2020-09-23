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

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	if rs := r.Conn.Create(&user); rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (r *userRepository) UserAlreadyExist(tx context.Context, email string) (bool, error) {
	var user domain.User
	if rs := r.Conn.Where("email = ?", email).Find(&user); rs.Error != nil {
		// ここでrecode not fountの時のエラーを返している。
		// todo: recode not found以外のエラーの判別したい
		return false, nil
	}
	if user.ID != 0 {
		return true, nil
	}
	return false, nil
}

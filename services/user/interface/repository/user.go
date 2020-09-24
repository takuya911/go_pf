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

func (r *userRepository) UpdateUser(ctx context.Context, formUser *domain.User) (*domain.User, *domain.User, error) {
	var aUser domain.User
	var bUser domain.User

	// transaction
	err := r.Conn.Transaction(func(tx *gorm.DB) error {
		// update after
		if rs := tx.First(&aUser, formUser.ID); rs.Error != nil {
			return rs.Error
		}
		// update
		if rs := tx.Save(&formUser); rs.Error != nil {
			return rs.Error
		}
		// update before
		if rs := tx.First(&bUser, formUser.ID); rs.Error != nil {
			return rs.Error
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return &aUser, &bUser, nil
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

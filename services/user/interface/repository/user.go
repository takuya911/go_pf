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

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	getUserSQL := "SELECT * FROM users WHERE id = ? AND deleted_at is null"
	if rs := r.Conn.Raw(getUserSQL, id).Scan(&user); rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	getUserSQL := "SELECT * FROM users WHERE email = ? AND deleted_at is null"
	if rs := r.Conn.Raw(getUserSQL, email).Scan(&user); rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, u *domain.User) error {
	if rs := r.Conn.Create(&u); rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, formUser *domain.User) (*domain.User, error) {
	var bUser domain.User

	// transaction
	err := r.Conn.Transaction(func(tx *gorm.DB) error {
		// update
		if rs := tx.Save(&formUser); rs.Error != nil {
			return rs.Error
		}
		// update before
		getUserSQL := "SELECT * FROM users WHERE id = ? AND deleted_at is null"
		if rs := r.Conn.Raw(getUserSQL, formUser.ID).Scan(&bUser); rs.Error != nil {
			return rs.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bUser, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, ID int64) (bool, error) {
	deleteUserSQL := "UPDATE users SET deleted_at = now() WHERE id = ?"
	if rs := r.Conn.Exec(deleteUserSQL, ID); rs.Error != nil {
		return false, rs.Error
	}
	return true, nil
}

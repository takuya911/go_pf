package domain

import (
	"time"
)

// User struct
type User struct {
	ID              int64      `json:"user_id" gorm:"primary_key"`
	Name            string     `json:"name" validate:"required"`
	Email           string     `json:"email" validate:"required,email"`
	Password        string     `json:"password" validate:"min=6,max=75"`
	TelephoneNumber string     `json:"telephone_number" validate:"required"`
	Gender          int64      `json:"gender" validate:"min=1,max=3"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

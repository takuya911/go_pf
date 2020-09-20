package domain

import "github.com/golang/protobuf/ptypes/timestamp"

// User struct
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt *timestamp.Timestamp
	UpdatedAt *timestamp.Timestamp
	DeletedAt *timestamp.Timestamp
}

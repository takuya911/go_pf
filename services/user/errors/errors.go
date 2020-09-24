package errors

import "errors"

var (
	PasswordFaultError = errors.New("password is incorrect")
	UserAlreadyExists  = errors.New("user is already exist, please register by other email")
	UserDoesNotExists  = errors.New("user does not exists")
)

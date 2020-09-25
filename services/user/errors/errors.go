package errors

import "errors"

var (
	PasswordFaultError = errors.New("password is incorrect")
	EmailAlreadyUsed   = errors.New("This email is already used, please register by other email")
)

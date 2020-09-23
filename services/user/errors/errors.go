package errors

import "errors"

var (
	BadRequestError   = errors.New("bad request")
	UserAlreadyExists = errors.New("user is already exist, please register by other email")
)

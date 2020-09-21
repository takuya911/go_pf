package errors

import "errors"

var (
	// BadRequestError ...
	BadRequestError = errors.New("無効な構文のため、サーバーはリクエストを理解できませんでした。")
)

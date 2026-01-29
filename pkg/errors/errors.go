package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user id not found in context")
	ErrInvalidUserIdType = errors.New("invalid user id type in context")
	ErrFileNotFound      = errors.New("file not found")
	ErrPermissionDenied  = errors.New("permission denied")
)

package errors

import "errors"

var (
	ErrUnimplemented = errors.New("unimplemented")
	ErrInvalidInput  = errors.New("invalid input")
	ErrCodecNotFound = errors.New("codec not found")
)

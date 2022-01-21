package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("your requested item is not found")
	ErrConflict            = errors.New("your item already exist")
	ErrExpired             = errors.New("your item already expired")
	ErrBadParamInput       = errors.New("given param is not valid")
	ErrUnAuthorized        = errors.New("your credential is not valid")
	ErrRedirect            = errors.New("your requested item redirect to another endpoint")
)

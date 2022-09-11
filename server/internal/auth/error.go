package auth

import "errors"

var (
	ErrInvalidCredentials  = errors.New(`invalid credentials`)
	ErrTokenIsNotPresented = errors.New(`token is not presented`)
	ErrInvalidAccessToken  = errors.New(`invalid access token`)
)

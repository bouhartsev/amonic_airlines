package errdomain

import (
	"errors"
	"fmt"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
	Details any    `json:"details,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

var (
	_ error = &ErrorResponse{}
)

var (
	InvalidJSONError = &ErrorResponse{
		Message: "Invalid JSON provided.",
		Type:    InvalidRequestType,
		Code:    "invalid_json",
	}

	InvalidCredentialsError = &ErrorResponse{
		Message: "Invalid credentials.",
		Type:    InvalidRequestType,
		Code:    "invalid_credentials",
	}

	NoActiveLoginsDetectedError = &ErrorResponse{
		Message: "No active logins detected.",
		Type:    InvalidRequestType,
		Code:    "no_active_logins",
	}
)

func NewInternalError(msg string) *ErrorResponse {
	fmt.Println("ERROR: " + msg)

	return &ErrorResponse{
		Message: msg,
		Type:    InternalType,
		Code:    "internal",
	}
}

var (
	ErrTokenIsNotPresented = errors.New(`token is not presented`)
	ErrInvalidAccessToken  = errors.New(`invalid access token`)
)

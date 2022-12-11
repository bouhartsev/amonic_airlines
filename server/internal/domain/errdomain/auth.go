package errdomain

var (
	InvalidAuthTokenError = ErrorResponse{
		Message: "Invalid auth token.",
		Type:    UnauthorizedType,
		Code:    "token:invalid",
	}

	AuthTokenIsNotPresentedError = ErrorResponse{
		Message: "Auth token is not presented.",
		Type:    UnauthorizedType,
		Code:    "token:empty",
	}
)

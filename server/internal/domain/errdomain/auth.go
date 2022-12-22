package errdomain

var (
	InvalidAuthTokenError = ErrorResponse{
		Message: "Invalid auth token.",
		Type:    UnauthorizedType,
		Code:    "token:invalid",
	}

	AuthTokenExpiredError = ErrorResponse{
		Message: "Auth token is expired.",
		Type:    UnauthorizedType,
		Code:    "token:expired",
	}

	AuthTokenIsNotPresentedError = ErrorResponse{
		Message: "Auth token is not presented.",
		Type:    UnauthorizedType,
		Code:    "token:empty",
	}
)

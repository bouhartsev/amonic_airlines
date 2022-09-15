package errdomain

var (
	UserDisabledError = &ErrorResponse{
		Message: "User disabled.",
		Type:    ObjectDisabledType,
		Code:    "user:disabled",
	}

	UserNotFoundError = &ErrorResponse{
		Message: "User not found.",
		Type:    ObjectNotFoundType,
		Code:    "user:not_found",
	}

	EmailAlreadyTakenError = &ErrorResponse{
		Message: "Email already taken.",
		Type:    ObjectDuplicateType,
		Code:    "user.email:already_taken",
	}
)

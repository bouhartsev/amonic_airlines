package errdomain

type ErrorResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

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
)

func NewInternalError(msg string) *ErrorResponse {
	return &ErrorResponse{
		Message: msg,
		Type:    InternalType,
		Code:    "internal",
	}
}

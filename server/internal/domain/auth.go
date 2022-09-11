package domain

type AuthSignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

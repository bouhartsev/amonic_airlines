package domain

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	jwt.RegisteredClaims
	RawToken string `json:"-"`
	User     *User  `json:"user"`
}

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type ReportLastLogoutErrorRequest struct {
	Reason        string `json:"reason"`
	SystemCrash   bool   `json:"systemCrash"`
	SoftwareCrash bool   `json:"softwareCrash"`
}

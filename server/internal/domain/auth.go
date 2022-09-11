package domain

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	jwt.RegisteredClaims
	User *User
}

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

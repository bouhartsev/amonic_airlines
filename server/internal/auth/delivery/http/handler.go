package http

import (
	"context"
	"database/sql"
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/golang-jwt/jwt"
)

type repository struct {
	conn *sql.DB
}

type AuthClaims struct {
	jwt.StandardClaims
}

func NewAuthRepository(conn *sql.DB) auth.Repository {
	return &repository{conn}
}

func (r *repository) SignIn(ctx context.Context, request *domain.AuthSignInRequest) {
}

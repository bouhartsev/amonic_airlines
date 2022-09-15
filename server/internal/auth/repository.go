package auth

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

type Repository interface {
	SignIn(ctx context.Context, request *domain.SignInRequest) (*domain.User, error)
}

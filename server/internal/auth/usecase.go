package auth

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

type UseCase interface {
	SignIn(ctx context.Context, request *domain.SignInRequest) (*string, error)
}

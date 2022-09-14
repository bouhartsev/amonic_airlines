package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (u *useCase) SignIn(ctx context.Context, request *domain.SignInRequest) (*string, error) {
	user, err := u.repo.SignIn(ctx, request)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errdomain.InvalidCredentialsError
		}

		return nil, errdomain.NewInternalError(err.Error())
	}

	claims := domain.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), // 30 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: user,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SigningString()

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return &token, nil
}

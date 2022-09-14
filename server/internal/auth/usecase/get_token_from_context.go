package usecase

import (
	"context"
	"fmt"
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/golang-jwt/jwt/v4"
)

func (u *useCase) GetTokenFromContext(ctx context.Context) (*domain.AuthClaims, error) {
	ctxToken := ctx.Value("token")

	if ctxToken == nil {
		return nil, auth.ErrTokenIsNotPresented
	}

	accessToken, ok := ctxToken.(string)

	if !ok {
		return nil, auth.ErrInvalidAccessToken
	}

	token, err := jwt.ParseWithClaims(accessToken, &domain.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(u.cfg.TokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, auth.ErrInvalidAccessToken
}

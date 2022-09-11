package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

type useCase struct {
	repo   auth.Repository
	logger *zap.Logger
	cfg    *config.Config
}

func NewAuthUseCase(repo auth.Repository, logger *zap.Logger, cfg *config.Config) auth.UseCase {
	return &useCase{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (u *useCase) SignIn(ctx context.Context, request *domain.SignInRequest) (*string, error) {
	user, err := u.repo.SignIn(ctx, request)

	if err != nil {
		if errors.Is(err, errdomain.ErrUserNotFound) {
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

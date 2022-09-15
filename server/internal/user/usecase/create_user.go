package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (u *useCase) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error) {
	found, err := u.repo.GetUserByEmail(ctx, request.Email)

	if err == nil && found != nil {
		return nil, errdomain.EmailAlreadyTakenError
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errdomain.NewInternalError(err.Error())
	}

	user, err := u.repo.CreateUser(ctx, request)

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return user, nil
}

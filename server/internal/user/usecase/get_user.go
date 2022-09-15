package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (u *useCase) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error) {
	user, err := u.repo.GetUserById(ctx, request.UserId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errdomain.UserNotFoundError
		}

		return nil, errdomain.NewInternalError(err.Error())
	}

	return user, nil
}

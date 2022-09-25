package usecase

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (u *useCase) GetUsers(ctx context.Context, request *domain.GetUsersRequest) ([]*domain.User, error) {

	// TODO: check office id
	//if request.OfficeId != nil {
	//
	//}

	users, err := u.repo.GetUsers(ctx, request)

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return users, err
}

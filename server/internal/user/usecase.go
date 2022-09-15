package user

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

type UseCase interface {
	CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error)
	GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error)
	//GetUsers(ctx context.Context, request *domain.GetUsersRequest) ([]*domain.User, error)
	//UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.User, error)
}

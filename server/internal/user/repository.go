package user

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error)
	GetUserById(ctx context.Context, userId int) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	//GetUsers(ctx context.Context, request *domain.GetUsersRequest) ([]*domain.User, error)
	//UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.User, error)
}

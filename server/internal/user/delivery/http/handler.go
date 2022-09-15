package http

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/user"
)

type handler struct {
	useCase user.UseCase
}

func NewUserHandler(uc user.UseCase) *handler {
	return &handler{useCase: uc}
}

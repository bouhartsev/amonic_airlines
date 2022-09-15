package http

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
)

type handler struct {
	useCase auth.UseCase
}

func NewAuthHandler(uc auth.UseCase) *handler {
	return &handler{useCase: uc}
}

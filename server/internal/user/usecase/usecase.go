package usecase

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/config"
	"github.com/bouhartsev/amonic_airlines/server/internal/user"
	"go.uber.org/zap"
)

type useCase struct {
	repo   user.Repository
	logger *zap.Logger
	cfg    *config.Config
}

func NewUserUseCase(repo user.Repository, logger *zap.Logger, cfg *config.Config) user.UseCase {
	return &useCase{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

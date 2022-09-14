package usecase

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/config"
	"go.uber.org/zap"
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

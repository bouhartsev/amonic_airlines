package server

import (
	"database/sql"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/config"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/database"
	"go.uber.org/zap"
)

type server struct {
	logger *zap.Logger
	cfg    *config.Config
	db     *sql.DB
}

func New(l *zap.Logger, c *config.Config) (*server, error) {
	s := &server{
		logger: l,
		cfg:    c,
	}

	conn, err := database.NewMySQLConnection(s.cfg.DatabaseConnection)

	if err != nil {
		return nil, err
	}

	s.db = conn

	return s, nil
}

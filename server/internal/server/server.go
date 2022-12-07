package server

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/bouhartsev/amonic_airlines/server/internal/core"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/config"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/database"
)

type Server struct {
	logger *zap.Logger
	cfg    *config.Config
	db     *sql.DB
	core   *core.Core
}

func New(l *zap.Logger, c *config.Config) (*Server, error) {
	s := &Server{
		logger: l,
		cfg:    c,
	}

	conn, err := database.NewMySQLConnection(s.cfg.DatabaseConnection)

	if err != nil {
		return nil, err
	}

	l.Info("Connected to database", zap.String("conn", c.DatabaseConnection))

	s.db = conn
	s.core = core.NewCore(l, s.db, c)

	return s, nil
}

func (s *Server) Run() error {
	router := s.initRoutes()

	httpServer := &http.Server{
		Addr:    s.cfg.AppPort,
		Handler: router,
	}

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("failed to listen and serve", zap.Error(err), zap.String("address", httpServer.Addr))
			quit <- os.Interrupt
		}
	}()

	s.logger.Info("Running the Server", zap.String("address", httpServer.Addr))

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return func() error {
		s.logger.Info("shutting down the Server...")

		err := s.db.Close()

		if err != nil {
			s.logger.Error(err.Error())
			return err
		}

		s.logger.Info("database had shut down")

		err = httpServer.Shutdown(ctx)

		if err != nil {
			s.logger.Error(err.Error())
			return err
		}

		s.logger.Info("Server had shut down successfully")

		return nil
	}()
}

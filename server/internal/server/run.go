package server

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func (s *server) Run() error {
	router := s.initRoutes()

	httpServer := &http.Server{
		Addr:         s.cfg.AppPort,
		Handler:      router,
		ReadTimeout:  1 << 20,
		WriteTimeout: 1 << 20,
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

	s.logger.Info("Running the server", zap.String("address", httpServer.Addr))

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return func() error {
		s.logger.Info("shutting down the server...")

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

		s.logger.Info("server had shut down successfully")

		return nil
	}()
}

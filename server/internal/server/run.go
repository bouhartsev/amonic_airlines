package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func (s *server) Run() error {
	router := gin.Default()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	router.GET("/api/heartbeat", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "Hello, world!")
	})

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

	return httpServer.Shutdown(ctx)
}

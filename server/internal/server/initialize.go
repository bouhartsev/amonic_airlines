package server

import (
	authHandler "github.com/bouhartsev/amonic_airlines/server/internal/auth/delivery/http"
	authRepository "github.com/bouhartsev/amonic_airlines/server/internal/auth/repository"
	authUseCase "github.com/bouhartsev/amonic_airlines/server/internal/auth/usecase"
	"github.com/gin-gonic/gin"
)

func (s *server) initRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Logger(),
	)

	api := router.Group(`/api`)

	authH := authHandler.NewAuthHandler(authUseCase.NewAuthUseCase(authRepository.NewAuthRepository(s.db), s.logger, s.cfg))

	handlers := []EndpointRegistrar{
		authH,
	}

	for _, h := range handlers {
		h.RegisterEndpoints(api)
	}

	return router
}

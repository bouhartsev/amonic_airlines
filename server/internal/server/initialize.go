package server

import (
	authHandler "github.com/bouhartsev/amonic_airlines/server/internal/auth/delivery/http"
	authRepository "github.com/bouhartsev/amonic_airlines/server/internal/auth/repository"
	authUseCase "github.com/bouhartsev/amonic_airlines/server/internal/auth/usecase"
	userHandler "github.com/bouhartsev/amonic_airlines/server/internal/user/delivery/http"
	userRepository "github.com/bouhartsev/amonic_airlines/server/internal/user/repository"
	userUseCase "github.com/bouhartsev/amonic_airlines/server/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

func (s *server) initRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Logger(),
	)

	api := router.Group(`/api`)

	handlers := []EndpointRegistrar{
		authHandler.NewAuthHandler(authUseCase.NewAuthUseCase(authRepository.NewAuthRepository(s.db), s.logger, s.cfg)),
		userHandler.NewUserHandler(userUseCase.NewUserUseCase(userRepository.NewUserRepository(s.db), s.logger, s.cfg)),
	}

	for _, h := range handlers {
		h.RegisterEndpoints(api)
	}

	return router
}

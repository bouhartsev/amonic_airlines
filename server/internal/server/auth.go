package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// SignIn godoc
// @Summary Авторизирует пользователя.
// @Description Возвращает JWT токен при успешной авторизации.
// @Description
// @Description Cписок возможных кодов ошибок:
// @Description * `invalid_credentials` - Неверный логин или пароль.
// @Description * `user:disabled` - Пользователь заблокирован.
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body domain.SignInRequest true "JSON input"
// @Success 200 {object} domain.SignInResponse
// @Failure 404 {object} errdomain.ErrorResponse
// @Failure 409 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/auth/sign-in [post]
func (s *Server) SignIn(c *gin.Context) {
	input := new(domain.SignInRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	response, err := s.core.SignIn(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// SignOut godoc
// @Summary Позволяет пользователю выйти из системы.
// @Description Cписок возможных кодов ошибок:
// @Description * `no_active_logins` - Попытка выйти из системы, когда пользователь уже вышел.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/auth/sign-out [post]
func (s *Server) SignOut(c *gin.Context) {
	err := s.core.SignOut(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// ReportLastLogoutError godoc
// @Summary Отправляет репорт об ошибке выхода пользователя из системы.
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body domain.ReportLastLogoutErrorRequest true "JSON input"
// @Success 200
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/auth/report [post]
func (s *Server) ReportLastLogoutError(c *gin.Context) {
	input := new(domain.ReportLastLogoutErrorRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	err := s.core.ReportLastLogoutError(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

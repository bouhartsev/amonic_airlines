package http

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser godoc
// @Summary Создает пользователя.
// @Description Создает пользователя. При дубликате email возвращает ошибку с кодом `user.email:already_taken`.
// @Tags Users
// @Accept json
// @Produce json
// @Param input body domain.CreateUserRequest true "JSON input"
// @Success 201 {object} domain.User
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users [post]
func (h *handler) CreateUser(c *gin.Context) {
	input := new(domain.CreateUserRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	user, err := h.useCase.CreateUser(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

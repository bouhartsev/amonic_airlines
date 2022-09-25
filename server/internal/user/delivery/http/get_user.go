package http

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Возвращает информацию о пользователе.
// @Description Если пользователь не найден, вернет ошибку с кодом `user:not_found`.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "Идентификатор пользователя"
// @Failure 404 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users/{user_id} [get]
func (h *handler) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		delivery.ErrorResponse(c, errdomain.UserNotFoundError)
		return
	}

	user, err := h.useCase.GetUser(c.Request.Context(), &domain.GetUserRequest{UserId: userId})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

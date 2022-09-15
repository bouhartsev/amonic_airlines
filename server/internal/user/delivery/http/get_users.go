package http

import (
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handler) GetUsers(c *gin.Context) {
	var officeId *int

	id, err := strconv.Atoi(c.Query("officeId"))

	if err == nil {
		officeId = &id
	}

	user, err := h.useCase.GetUsers(c.Request.Context(), &domain.GetUsersRequest{OfficeId: officeId})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

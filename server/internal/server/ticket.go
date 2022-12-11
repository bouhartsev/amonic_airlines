package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// AddTicket godoc
// @Summary Добавляет билет.
// @Description Поле return опционально.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param input body domain.AddTicketRequest true "JSON input"
// @Success 201
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users [post]
func (s *Server) AddTicket(c *gin.Context) {
	input := new(domain.AddTicketRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	err := s.core.AddTicket(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

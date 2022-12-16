package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// GetDetailedReport godoc
// @Summary Возвращает информацию о кратком обзоре из `5.5`.
// @Tags Reports
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetDetailedReportResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/report [get]
func (s *Server) GetDetailedReport(c *gin.Context) {
	response, err := s.core.GetDetailedReport(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// GetAmenities godoc
// @Summary Возвращает список сервисов.
// @Tags Amenities
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetAmenitiesResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/amenities [get]
func (s *Server) GetAmenities(c *gin.Context) {
	response, err := s.core.GetAmenities(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

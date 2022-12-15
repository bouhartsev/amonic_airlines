package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
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

// GetAmenitiesBriefReport godoc
// @Summary Возвращает краткий отчет о сервисах.
// @Tags Amenities
// @Accept json
// @Produce json
// @Param schedule_id query int false "Фильтрация по полёту"
// @Param from query int false "Начало временного отрезка фильтрации формата 2010-10-31"
// @Param to query int false "Конец временного отрезка фильтрации формата 2010-10-31"
// @Success 200 {object} domain.GetAmenitiesBriefReportResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/amenities/reports/brief [get]
func (s *Server) GetAmenitiesBriefReport(c *gin.Context) {
	var (
		scheduleId       *int
		dateFrom, dateTo *string
	)

	if val, err := strconv.Atoi(c.Query("schedule_id")); err == nil {
		scheduleId = &val
	}
	if val := c.Query("from"); val != "" {
		dateFrom = &val
	}
	if val := c.Query("to"); val != "" {
		dateTo = &val
	}

	response, err := s.core.GetAmenitiesBriefReport(c.Request.Context(), &domain.GetAmenitiesBriefReportRequest{
		DateFrom:   dateFrom,
		DateTo:     dateTo,
		ScheduleId: scheduleId,
	})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

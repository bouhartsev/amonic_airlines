package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// GetSchedules godoc
// @Summary Возвращает список расписаний полетов(schedules).
// @Tags Schedules
// @Accept json
// @Produce json
// @Param from query string false "Имя аэропорта, в который идет отправление"
// @Param to query string false "Имя аэропорта, из коготорого идет отправление"
// @Param sort_by query string false "Сортировка. Возможные значения: `datetime`, `price`, `confirmed`, `unconfirmed`. По умолчанию `datetime`."
// @Param outbound query string false "Фильтр по дате вылета"
// @Param flightNumber query string false "Фильтр по номету полета"
// @Success 200 {object} domain.GetSchedulesResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/schedules [get]
func (s *Server) GetSchedules(c *gin.Context) {
	var (
		from         *string
		to           *string
		sortBy       *string
		outbound     *string
		flightNumber *int
	)

	if val := c.Query("from"); val != "" {
		from = &val
	}
	if val := c.Query("to"); val != "" {
		to = &val
	}
	if val := c.Query("sort_by"); val != "" {
		sortBy = &val
	}
	if val := c.Query("outbound"); val != "" {
		outbound = &val
	}
	if val := c.Query("flightNumber"); val != "" {
		if v, err := strconv.Atoi(val); err == nil {
			flightNumber = &v
		}
	}

	schedules, err := s.core.GetSchedules(c.Request.Context(), &domain.GetSchedulesRequest{
		From:         from,
		To:           to,
		SortBy:       sortBy,
		Outbound:     outbound,
		FlightNumber: flightNumber,
	})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// SwitchScheduleStatus godoc
// @Summary Переключает флаг confirmed.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule_id path int true "Идентификатор расписания"
// @Success 200
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/schedules/{schedule_id}/switch-status [post]
func (s *Server) SwitchScheduleStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("schedule_id"))

	err := s.core.SwitchScheduleStatus(c.Request.Context(), &domain.SwitchScheduleStatusRequest{ScheduleId: id})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// UpdateSchedule godoc
// @Summary Обновляет расписание.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule_id path int true "Идентификатор расписания"
// @Param input body domain.UpdateScheduleRequest true "JSON input"
// @Success 200
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/schedules/{schedule_id} [patch]
func (s *Server) UpdateSchedule(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("schedule_id"))

	input := new(domain.UpdateScheduleRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	input.ScheduleId = id

	err := s.core.UpdateSchedule(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// UpdateSchedulesFromFile godoc
// @Summary Получает данные для изменений из файла, парсит и применяет их.
// @Tags Schedules
// @Accept multipart/form-data
// @Produce json
// @Param input formData file true "Файл со списком изменений"
// @Success 200 {object} domain.UpdateSchedulesFromFileResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/schedules/upload [post]
func (s *Server) UpdateSchedulesFromFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &errdomain.ErrorResponse{
			Message: err.Error(),
			Type:    errdomain.InvalidRequestType,
		})
		return
	}

	file, err := fileHeader.Open()
	defer file.Close()

	response, err := s.core.UpdateSchedulesFromFile(c.Request.Context(), file)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

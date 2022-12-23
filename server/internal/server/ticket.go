package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// AddTicket godoc
// @Summary Добавляет билет.
// @Description Поле return опционально.
// @Tags  Tickets
// @Accept json
// @Produce json
// @Param input body domain.AddTicketRequest true "JSON input"
// @Success 201
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/tickets [post]
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

// GetTickets godoc
// @Summary Возвращает список билетов.
// @Description Список может содержать до 100 элементов. Главным элементом поиска является параметр `booking_reference`.
// @Description У билетов "туда и обратно" он одинаковый.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param user_id query int false "Фильтрация по пользователю, которому принадлежат билеты"
// @Param schedule_id query int false "Фильтрация по полету, которому принадлежат билеты"
// @Param booking_reference query int false "Фильтрация по брони, которой принадлежат билеты"
// @Success 200 {object} domain.GetTicketsResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/tickets [get]
func (s *Server) GetTickets(c *gin.Context) {
	var (
		userId           *int
		scheduleId       *int
		bookingReference *string
	)

	if val, err := strconv.Atoi(c.Query("user_id")); err == nil {
		userId = &val
	}
	if val, err := strconv.Atoi(c.Query("schedule_id")); err == nil {
		scheduleId = &val
	}
	if val := c.Query("booking_reference"); val != "" {
		bookingReference = &val
	}

	users, err := s.core.GetTickets(c.Request.Context(), &domain.GetTicketsRequest{
		UserId:           userId,
		ScheduleId:       scheduleId,
		BookingReference: bookingReference,
	})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetTicketAmenities godoc
// @Summary Возвращает список выбранных сервисов для указанного билета.
// @Tags Tickets
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetTicketAmenitiesResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/tickets/{ticket_id}/amenities [get]
func (s *Server) GetTicketAmenities(c *gin.Context) {
	ticketId, _ := strconv.Atoi(c.Param("ticket_id"))
	response, err := s.core.GetTicketAmenities(c.Request.Context(), ticketId)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// RemoveTicketAmenities godoc
// @Summary Удаляет список сервисов для указанного билета.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param input body domain.RemoveTicketAmenitiesRequest true "JSON input"
// @Success 200
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/tickets/{ticket_id}/amenities [delete]
func (s *Server) RemoveTicketAmenities(c *gin.Context) {
	input := new(domain.RemoveTicketAmenitiesRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}
	input.TicketId, _ = strconv.Atoi(c.Param("ticket_id"))

	err := s.core.RemoveTicketAmenities(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// AddTicketAmenities godoc
// @Summary Добавляет список сервисов для указанного билета.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param input body domain.AddTicketAmenitiesRequest true "JSON input"
// @Success 200
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/tickets/{ticket_id}/amenities [post]
func (s *Server) AddTicketAmenities(c *gin.Context) {
	input := new(domain.AddTicketAmenitiesRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}
	input.TicketId, _ = strconv.Atoi(c.Param("ticket_id"))

	err := s.core.AddTicketAmenities(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

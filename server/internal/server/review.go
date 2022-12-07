package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// AddReview godoc
// @Summary Добавляет отзыв.
// @Description У поля `gender` 0 - Male, 1 - Female.
// @Tags Reviews
// @Accept json
// @Produce json
// @Param input body domain.AddReviewRequest true "JSON input"
// @Success 201
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/reviews [post]
func (s *Server) AddReview(c *gin.Context) {
	input := new(domain.AddReviewRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	err := s.core.AddReview(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

// GetReviewsBrief godoc
// @Summary Возвращает суммарный отчет.
// @Description Поля `to` и `from` в query обязательны.
// @Tags Reviews
// @Accept json
// @Produce json
// @Param to query string true "Дата, представленная начальной границей выборки"
// @Param from query string true "Дата, представленная конечной границей выборки"
// @Success 200 {object} domain.GetBriefReviewsResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/reviews/brief [get]
func (s *Server) GetReviewsBrief(c *gin.Context) {
	response, err := s.core.GetReviewsBrief(c.Request.Context(), &domain.GetBriefReviewsRequest{
		From: c.Query("from"),
		To:   c.Query("to"),
	})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

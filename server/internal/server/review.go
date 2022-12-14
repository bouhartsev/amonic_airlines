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
// @Description
// @Description `qN` - ответ от `0` до `7`.
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

// GetBriefReviews godoc
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
func (s *Server) GetBriefReviews(c *gin.Context) {
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

// GetDetailedReviews godoc
// @Summary Возвращает подробный суммарный отчет.
// @Description Поля `to` и `from` в query обязательны.
// @Description Дата формата `2022-01-27 07:30:54`.
// @Description
// @Description `q1`, `q2`, `q3`, `q4` - Объекты, содержащие информацию четырех вопросов.
// @Description `1...7` - Информация о количестве ответов(горизонтальная линия).
// @Tags Reviews
// @Accept json
// @Produce json
// @Param to query string true "Дата, представленная начальной границей выборки"
// @Param from query string true "Дата, представленная конечной границей выборки"
// @Success 200 {object} domain.GetDetailedReviewsResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/reviews/detailed [get]
func (s *Server) GetDetailedReviews(c *gin.Context) {
	response, err := s.core.GetDetailedReviews(c.Request.Context(), &domain.GetDetailedReviewsRequest{
		From: c.Query("from"),
		To:   c.Query("to"),
	})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

package http

import (
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) SignIn(c *gin.Context) {
	input := new(domain.SignInRequest)

	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), input)

	if err != nil {
		if errors.Is(err, errdomain.InvalidCredentialsError) {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &domain.SignInResponse{Token: *token})
}

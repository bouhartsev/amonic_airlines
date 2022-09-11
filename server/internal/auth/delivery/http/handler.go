package http

import (
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	useCase auth.UseCase
}

func NewAuthHandler(uc auth.UseCase) *handler {
	return &handler{useCase: uc}
}

func (h *handler) SignIn(c *gin.Context) {
	input := new(domain.SignInRequest)

	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), input)

	// TODO: check why it doesn't return json properly
	if err != nil {
		if errors.Is(err, errdomain.InvalidCredentialsError) {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err})
			utils.ErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &domain.SignInResponse{Token: *token})
}

package http

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterEndpoints(router gin.IRouter) {
	endpoint := router.Group(`/auth`)

	// TODO: sign-out
	endpoint.POST(`/sign-in`, h.SignIn)
}

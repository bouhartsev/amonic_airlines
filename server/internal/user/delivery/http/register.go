package http

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterEndpoints(router gin.IRouter) {
	endpoint := router.Group(`/users`)

	// TODO: sign-out
	endpoint.POST(`/`, h.CreateUser)
	endpoint.GET(`/`, h.GetUsers)
	endpoint.GET(`/:user_id`, h.GetUser)
}

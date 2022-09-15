package server

import "github.com/gin-gonic/gin"

type EndpointRegistrar interface {
	RegisterEndpoints(engine gin.IRouter)
}

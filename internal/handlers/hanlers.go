package handlers

import (
	"github.com/fishmanDK/internal/service"
	"github.com/gin-gonic/gin"
)


type Handlers struct{
	Service *service.Service
}


func NewHandlers(serv *service.Service) *Handlers{
	return &Handlers{
		Service: serv,
	}
}


func(h *Handlers) InitRouts() *gin.Engine{
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.authentication)
		auth.POST("/sign-up", h.createUser)
	}

	return router
}
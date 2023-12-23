package h_http

import (
	"github.com/fishmanDK/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handlers struct {
	Service *service.Service
}

func NewHandlers(serv *service.Service) *Handlers {
	return &Handlers{
		Service: serv,
	}
}

func (h *Handlers) InitRouts(logger *slog.Logger) *gin.Engine {
	router := gin.Default()
	router.Use(LoggerMiddleware(logger))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.authentication)
		auth.POST("/sign-up", h.createUser)
	}

	return router
}

package handler

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *handler {
	return &handler{services}
}

func (h *handler) InitRouter() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/token", h.Token)
		auth.POST("/sign-in", h.SignIn)
	}

	return r
}

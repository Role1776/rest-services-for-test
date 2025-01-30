package handler

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) InitRouter() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := r.Group("/api", h.userIdentityMiddleware)
	{
		api.GET("/move", h.Get)
		api.POST("/move", h.Create)
	}

	return r
}

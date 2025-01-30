package handler

import (
	"app/internal/repository"

	"github.com/gin-gonic/gin"
)

type handler struct {
	repos *repository.Repository
}

func NewHandler(repo repository.Repository) *handler {
	return &handler{repos: &repo}
}

func (h *handler) InitRouter() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUpDB)
		auth.POST("/sign-in", h.SignInDB)
	}

	api := r.Group("/api")
	{
		api.POST("/move", h.GetDB)
		api.POST("/move/create", h.CreateDB)
	}

	return r
}

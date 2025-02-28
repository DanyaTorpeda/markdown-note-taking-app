package handler

import (
	"markdown-note/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	notes := router.Group("/notes")
	{
		notes.POST("/", h.createNote)
	}
	return router
}

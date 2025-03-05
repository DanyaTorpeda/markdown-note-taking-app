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

	api := router.Group("/api")
	notes := api.Group("/notes")
	{
		notes.POST("/", h.createNote)
		notes.GET("/:id", h.getById)
		notes.POST("/:id/attachments", h.createAttachments)
		notes.GET("/:id/uploads/:file_name", h.getAttachment)
	}
	return router
}

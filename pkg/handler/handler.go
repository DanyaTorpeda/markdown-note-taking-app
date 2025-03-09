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
		notes.PUT("/:id", h.updateNote)
		notes.DELETE("/:id", h.deleteNote)
		notes.POST("/:id/attachments", h.createAttachments)
		notes.DELETE("/:id/attachments/:file_id", h.deleteAttachment)
		notes.GET("/:id/uploads/:file_name", h.getAttachment)
		notes.GET("/:id/render", h.parsingHTML)
	}
	return router
}

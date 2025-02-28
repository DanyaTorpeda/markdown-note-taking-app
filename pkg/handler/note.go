package handler

import (
	"io"
	mdnote "markdown-note"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createNote(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := header.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input := mdnote.Note{
		Title:   header.Filename,
		Content: string(fileContent),
	}

	id, err := h.service.Note.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

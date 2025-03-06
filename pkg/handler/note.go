package handler

import (
	"io"
	mdnote "markdown-note"
	"net/http"
	"strconv"

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

func (h *Handler) getById(c *gin.Context) {
	val := c.Param("id")
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id data")
		return
	}

	id, err := strconv.Atoi(val)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	note, err := h.service.Note.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *Handler) deleteNote(c *gin.Context) {
	val := c.Param("id")
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id data")
		return
	}

	id, err := strconv.Atoi(val)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Note.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

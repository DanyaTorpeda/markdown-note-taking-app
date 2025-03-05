package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createAttachments(c *gin.Context) {
	val := c.Param("id")
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	id, err := strconv.Atoi(val)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	headers, ok := form.File["files"]
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no files found")
		return
	}

	err = h.service.Attachment.Create(id, headers)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "no files found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "status": "ok"})
}

func (h *Handler) getAttachment(c *gin.Context) {
	noteId := c.Param("id")
	fileName := c.Param("file_name")
	absPath, err := filepath.Abs(fmt.Sprintf("uploads/%s/%s", noteId, fileName))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.File(absPath)
}

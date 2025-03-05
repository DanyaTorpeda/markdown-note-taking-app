package service

import (
	"fmt"
	"io"
	mdnote "markdown-note"
	"markdown-note/pkg/repository"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type AttachmentService struct {
	repo repository.Attachment
}

func NewAttachmentService(repo repository.Attachment) *AttachmentService {
	return &AttachmentService{repo: repo}
}

func (s *AttachmentService) Create(noteId int, headers []*multipart.FileHeader) error {
	err := os.MkdirAll(fmt.Sprintf("uploads/%d", noteId), os.ModePerm)
	if err != nil {
		return err
	}

	var attachments []mdnote.Attachment
	for _, header := range headers {
		file, err := header.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		body, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("uploads/%d/%s", noteId, header.Filename), body, 0644)
		if err != nil {
			return err
		}

		attachments = append(attachments, mdnote.Attachment{
			FileName: header.Filename,
			FilePath: filepath.Join("uploads", strconv.Itoa(noteId), header.Filename),
			FileType: http.DetectContentType(body),
			FileSize: header.Size,
		})
	}

	return s.repo.Create(noteId, attachments)
}

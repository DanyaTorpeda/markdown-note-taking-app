package service

import (
	mdnote "markdown-note"
	"markdown-note/pkg/repository"
	"mime/multipart"
)

type Note interface {
	Create(input mdnote.Note) (int, error)
	GetById(id int) (*mdnote.Note, error)
	Update(id int, input mdnote.Note) error
	Delete(id int) error
}

type Attachment interface {
	Create(noteId int, headers []*multipart.FileHeader) error
	Delete(noteId int, fileId int) error
}

type Service struct {
	Note
	Attachment
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Note:       NewNoteService(repository.Note),
		Attachment: NewAttachmentService(repository.Attachment),
	}
}

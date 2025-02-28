package service

import (
	mdnote "markdown-note"
	"markdown-note/pkg/repository"
)

type Note interface {
	Create(input mdnote.Note) (int, error)
}

type Service struct {
	Note
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Note: NewNoteService(repository.Note),
	}
}

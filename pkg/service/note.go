package service

import (
	mdnote "markdown-note"
	"markdown-note/pkg/repository"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(input mdnote.Note) (int, error) {
	return s.repo.Create(input)
}

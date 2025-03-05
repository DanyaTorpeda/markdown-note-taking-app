package service

import (
	"errors"
	mdnote "markdown-note"
	"markdown-note/pkg/repository"
	"strings"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(input mdnote.Note) (int, error) {
	if strings.TrimSpace(input.Content) == "" {
		return 0, errors.New("note content cannot be empty")
	}
	return s.repo.Create(input)
}

func (s *NoteService) GetById(id int) (*mdnote.Note, error) {
	return s.repo.GetById(id)
}

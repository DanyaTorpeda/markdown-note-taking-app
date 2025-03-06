package service

import (
	"errors"
	"fmt"
	mdnote "markdown-note"
	"markdown-note/pkg/repository"
	"os"
	"path/filepath"
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

func (s *NoteService) Delete(id int) error {
	dirPath, err := filepath.Abs(fmt.Sprintf("uploads/%d", id))
	if err != nil {
		return err
	}

	err = os.RemoveAll(dirPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete folder %s: %w", dirPath, err)
	}

	return s.repo.Delete(id)
}

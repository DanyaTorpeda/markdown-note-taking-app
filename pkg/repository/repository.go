package repository

import (
	mdnote "markdown-note"

	"github.com/jmoiron/sqlx"
)

type Note interface {
	Create(input mdnote.Note) (int, error)
}

type Repository struct {
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Note: NewNotePostgres(db),
	}
}

package repository

import (
	mdnote "markdown-note"

	"github.com/jmoiron/sqlx"
)

type Note interface {
	Create(input mdnote.Note) (int, error)
	GetById(id int) (*mdnote.Note, error)
}

type Attachment interface {
	Create(noteId int, attachments []mdnote.Attachment) error
}

type Repository struct {
	Note
	Attachment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Note:       NewNotePostgres(db),
		Attachment: NewAttachmentPostgres(db),
	}
}

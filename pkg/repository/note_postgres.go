package repository

import (
	"database/sql"
	"errors"
	"fmt"
	mdnote "markdown-note"

	"github.com/jmoiron/sqlx"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) Create(input mdnote.Note) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, content) values($1, $2) RETURNING id", notesTable)
	row := r.db.QueryRow(query, input.Title, input.Content)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NotePostgres) GetById(id int) (*mdnote.Note, error) {
	queryGetNote := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", notesTable)
	var note mdnote.Note
	err := r.db.Get(&note, queryGetNote, id)
	if err != nil {
		return nil, err
	}

	queryGetAttachments := fmt.Sprintf("SELECT * FROM %s WHERE note_id = $1", attachmentsTable)
	var attachments []mdnote.Attachment
	err = r.db.Select(&attachments, queryGetAttachments, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	note.Attachments = attachments
	return &note, nil
}

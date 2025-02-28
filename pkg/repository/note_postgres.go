package repository

import (
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

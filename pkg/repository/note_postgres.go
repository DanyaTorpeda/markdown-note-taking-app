package repository

import (
	"database/sql"
	"errors"
	"fmt"
	mdnote "markdown-note"
	"strings"

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

func (r *NotePostgres) Update(id int, input mdnote.Note) error {
	var placeholders []string
	var args []interface{}
	indexId := 1

	if input.Title != "" {
		placeholders = append(placeholders, fmt.Sprintf("title = $%d", indexId))
		args = append(args, input.Title)
		indexId++
	}

	if input.Content != "" {
		placeholders = append(placeholders, fmt.Sprintf("content = $%d", indexId))
		args = append(args, input.Content)
		indexId++
	}

	placeholders = append(placeholders, fmt.Sprintf("updated_at = $%d", indexId))
	args = append(args, "now()")
	indexId++

	params := strings.Join(placeholders, ", ")

	updateQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", notesTable, params, indexId)
	args = append(args, id)

	res, err := r.db.Exec(updateQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nothing updated")
	}

	return nil
}

func (r *NotePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", notesTable)
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nothing was deleted")
	}

	return nil
}

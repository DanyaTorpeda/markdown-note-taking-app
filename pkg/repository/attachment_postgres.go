package repository

import (
	"errors"
	"fmt"
	mdnote "markdown-note"
	"strings"

	"github.com/jmoiron/sqlx"
)

type AttachmentPostgres struct {
	db *sqlx.DB
}

func NewAttachmentPostgres(db *sqlx.DB) *AttachmentPostgres {
	return &AttachmentPostgres{db: db}
}

func (r *AttachmentPostgres) Create(noteId int, attachments []mdnote.Attachment) error {
	if len(attachments) == 0 {
		return errors.New("no files found")
	}

	query := fmt.Sprintf("INSERT INTO %s (note_id, file_name, file_path, file_type, file_size) VALUES ", attachmentsTable)
	values := []interface{}{}
	placeholders := []string{}

	for i, att := range attachments {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		values = append(values, noteId, att.FileName, att.FilePath, att.FileType, att.FileSize)
	}

	query += strings.Join(placeholders, ", ")
	_, err := r.db.Exec(query, values...)
	return err
}

func (r *AttachmentPostgres) Delete(noteId int, fileId int) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND note_id = $2 RETURNING file_name", attachmentsTable)
	row := r.db.QueryRow(query, fileId, noteId)
	var fileName string
	err := row.Scan(&fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

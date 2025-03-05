package mdnote

import "time"

type Note struct {
	ID          int          `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Content     string       `json:"content" db:"content"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	ID       int    `json:"id" db:"id"`
	NoteID   int    `json:"note_id" db:"note_id"`
	FileName string `json:"file_name" db:"file_name"`
	FilePath string `json:"file_path" db:"file_path"`
	FileType string `json:"file_type" db:"file_type"`
	FileSize int64  `json:"file_size" db:"file_size"`
}

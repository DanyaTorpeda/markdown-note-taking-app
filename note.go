package mdnote

import "time"

type Note struct {
	ID          int
	Title       string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Attachments []Attachment
}

type Attachment struct {
	ID       int
	NoteID   int
	FileName string
	FilePath string
	FileType string
	FileSize int64
}

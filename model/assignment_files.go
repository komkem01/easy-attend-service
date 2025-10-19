package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AssignmentFiles table structure
type AssignmentFiles struct {
	bun.BaseModel `bun:"table:assignment_files,alias:af"`

	ID           uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	AssignmentID uuid.UUID `json:"assignment_id" bun:"assignment_id,notnull,type:uuid"`
	FileName     string    `json:"file_name" bun:"file_name,notnull"`
	FilePath     string    `json:"file_path" bun:"file_path,notnull"`
	FileSize     int       `json:"file_size" bun:"file_size,notnull"`
	FileType     string    `json:"file_type" bun:"file_type,notnull"`
	UploadedBy   uuid.UUID `json:"uploaded_by" bun:"uploaded_by,notnull,type:uuid"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`

	// Relations
	Assignment *Assignments `json:"assignment,omitempty" bun:"rel:belongs-to,join:assignment_id=id"`
	Uploader   *Users       `json:"uploader,omitempty" bun:"rel:belongs-to,join:uploaded_by=id"`
}

// TableName returns the table name
func (af *AssignmentFiles) TableName() string {
	return "assignment_files"
}

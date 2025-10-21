package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// FileUploads table structure
type FileUploads struct {
	bun.BaseModel `bun:"table:file_uploads,alias:fu"`

	ID           uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OriginalName string     `json:"original_name" bun:"original_name,notnull"`
	StoredName   string     `json:"stored_name" bun:"stored_name,notnull"`
	FilePath     string     `json:"file_path" bun:"file_path,notnull"`
	FileSize     int64      `json:"file_size" bun:"file_size,notnull"`
	MimeType     string     `json:"mime_type" bun:"mime_type,notnull"`
	Category     string     `json:"category" bun:"category,notnull,type:file_category"`
	UploadedBy   uuid.UUID  `json:"uploaded_by" bun:"uploaded_by,notnull,type:uuid"`
	RelatedTable *string    `json:"related_table" bun:"related_table"`
	RelatedID    *uuid.UUID `json:"related_id" bun:"related_id,type:uuid"`
	IsPublic     bool       `json:"is_public" bun:"is_public,notnull,default:false"`
	Description  *string    `json:"description" bun:"description"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`
	CreatedAt    time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt    time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relationships
	Uploader *Users `json:"uploader,omitempty" bun:"rel:belongs-to,join:uploaded_by=id"`
}

// TableName returns the table name
func (fu *FileUploads) TableName() string {
	return "file_uploads"
}

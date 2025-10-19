package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Assignments table structure
type Assignments struct {
	bun.BaseModel `bun:"table:assignments,alias:a"`

	ID                  uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID         uuid.UUID  `json:"classroom_id" bun:"classroom_id,notnull,type:uuid"`
	Title               string     `json:"title" bun:"title,notnull"`
	Description         *string    `json:"description" bun:"description"`
	Instructions        *string    `json:"instructions" bun:"instructions"`
	AssignmentType      string     `json:"assignment_type" bun:"assignment_type,notnull,default:'homework',type:assignment_type"`
	DueDate             *time.Time `json:"due_date" bun:"due_date"`
	MaxScore            float64    `json:"max_score" bun:"max_score,default:100.00"`
	Weight              float64    `json:"weight" bun:"weight,default:1.00"`
	AllowLateSubmission bool       `json:"allow_late_submission" bun:"allow_late_submission,notnull,default:false"`
	LatePenaltyPercent  float64    `json:"late_penalty_percent" bun:"late_penalty_percent,default:0"`
	SubmissionFormat    string     `json:"submission_format" bun:"submission_format,notnull,default:'both',type:submission_format"`
	MaxFileSizeMB       int        `json:"max_file_size_mb" bun:"max_file_size_mb,default:10"`
	AllowedFileTypes    *string    `json:"allowed_file_types" bun:"allowed_file_types,type:jsonb"`
	IsPublished         bool       `json:"is_published" bun:"is_published,notnull,default:false"`
	Status              string     `json:"status" bun:"status,notnull,default:'draft',type:assignment_status"`
	CreatedBy           uuid.UUID  `json:"created_by" bun:"created_by,notnull,type:uuid"`
	CreatedAt           time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt           time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relations
	Classroom *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
	Creator   *Users      `json:"creator,omitempty" bun:"rel:belongs-to,join:created_by=id"`
}

// TableName returns the table name
func (a *Assignments) TableName() string {
	return "assignments"
}

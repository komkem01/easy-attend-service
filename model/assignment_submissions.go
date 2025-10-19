package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AssignmentSubmissions table structure
type AssignmentSubmissions struct {
	bun.BaseModel `bun:"table:assignment_submissions,alias:asub"`

	ID             uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	AssignmentID   uuid.UUID  `json:"assignment_id" bun:"assignment_id,notnull,type:uuid"`
	StudentID      uuid.UUID  `json:"student_id" bun:"student_id,notnull,type:uuid"`
	SubmissionText *string    `json:"submission_text" bun:"submission_text"`
	Status         string     `json:"status" bun:"status,notnull,default:'draft',type:submission_status"`
	SubmittedAt    *time.Time `json:"submitted_at" bun:"submitted_at"`
	IsLate         bool       `json:"is_late" bun:"is_late,notnull,default:false"`
	LateMinutes    int        `json:"late_minutes" bun:"late_minutes,default:0"`
	Score          *float64   `json:"score" bun:"score,type:numeric"`
	Feedback       *string    `json:"feedback" bun:"feedback"`
	GradedBy       *uuid.UUID `json:"graded_by" bun:"graded_by,type:uuid"`
	GradedAt       *time.Time `json:"graded_at" bun:"graded_at"`
	CreatedAt      time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt      time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relations
	Assignment *Assignments `json:"assignment,omitempty" bun:"rel:belongs-to,join:assignment_id=id"`
	Student    *Users       `json:"student,omitempty" bun:"rel:belongs-to,join:student_id=id"`
	Grader     *Users       `json:"grader,omitempty" bun:"rel:belongs-to,join:graded_by=id"`
}

// TableName returns the table name
func (asub *AssignmentSubmissions) TableName() string {
	return "assignment_submissions"
}

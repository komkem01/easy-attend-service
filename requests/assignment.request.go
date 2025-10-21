package requests

import (
	"time"

	"github.com/google/uuid"
)

// CreateAssignmentRequest for creating new assignment
type CreateAssignmentRequest struct {
	ClassroomID         uuid.UUID  `json:"classroom_id" validate:"required,uuid"`
	Title               string     `json:"title" validate:"required,min=2,max=200"`
	Description         *string    `json:"description" validate:"omitempty,max=2000"`
	Instructions        *string    `json:"instructions" validate:"omitempty,max=5000"`
	AssignmentType      *string    `json:"assignment_type" validate:"omitempty,oneof=homework quiz exam project lab_work"`
	DueDate             *time.Time `json:"due_date" validate:"omitempty"`
	MaxScore            *float64   `json:"max_score" validate:"omitempty,min=0"`
	Weight              *float64   `json:"weight" validate:"omitempty,min=0"`
	AllowLateSubmission *bool      `json:"allow_late_submission" validate:"omitempty"`
	LatePenaltyPercent  *float64   `json:"late_penalty_percent" validate:"omitempty,min=0,max=100"`
	SubmissionFormat    *string    `json:"submission_format" validate:"omitempty,oneof=text file both"`
	MaxFileSizeMB       *int       `json:"max_file_size_mb" validate:"omitempty,min=1,max=100"`
	AllowedFileTypes    *string    `json:"allowed_file_types" validate:"omitempty"`
	IsPublished         *bool      `json:"is_published" validate:"omitempty"`
}

// UpdateAssignmentRequest for updating assignment
type UpdateAssignmentRequest struct {
	Title               *string    `json:"title" validate:"omitempty,min=2,max=200"`
	Description         *string    `json:"description" validate:"omitempty,max=2000"`
	Instructions        *string    `json:"instructions" validate:"omitempty,max=5000"`
	AssignmentType      *string    `json:"assignment_type" validate:"omitempty,oneof=homework quiz exam project lab_work"`
	DueDate             *time.Time `json:"due_date" validate:"omitempty"`
	MaxScore            *float64   `json:"max_score" validate:"omitempty,min=0"`
	Weight              *float64   `json:"weight" validate:"omitempty,min=0"`
	AllowLateSubmission *bool      `json:"allow_late_submission" validate:"omitempty"`
	LatePenaltyPercent  *float64   `json:"late_penalty_percent" validate:"omitempty,min=0,max=100"`
	SubmissionFormat    *string    `json:"submission_format" validate:"omitempty,oneof=text file both"`
	MaxFileSizeMB       *int       `json:"max_file_size_mb" validate:"omitempty,min=1,max=100"`
	AllowedFileTypes    *string    `json:"allowed_file_types" validate:"omitempty"`
	IsPublished         *bool      `json:"is_published" validate:"omitempty"`
	Status              *string    `json:"status" validate:"omitempty,oneof=draft active completed archived"`
}

// AssignmentQueryRequest for filtering assignments
type AssignmentQueryRequest struct {
	Page           int        `json:"page" query:"page" validate:"omitempty,min=1"`
	Limit          int        `json:"limit" query:"limit" validate:"omitempty,min=1,max=100"`
	Search         *string    `json:"search" query:"search" validate:"omitempty,min=1,max=100"`
	ClassroomID    *uuid.UUID `json:"classroom_id" query:"classroom_id" validate:"omitempty,uuid"`
	CreatedBy      *uuid.UUID `json:"created_by" query:"created_by" validate:"omitempty,uuid"`
	AssignmentType *string    `json:"assignment_type" query:"assignment_type" validate:"omitempty,oneof=homework quiz exam project lab_work"`
	Status         *string    `json:"status" query:"status" validate:"omitempty,oneof=draft active completed archived"`
	IsPublished    *bool      `json:"is_published" query:"is_published" validate:"omitempty"`
	DueSoon        *bool      `json:"due_soon" query:"due_soon" validate:"omitempty"` // Due within 7 days
	Overdue        *bool      `json:"overdue" query:"overdue" validate:"omitempty"`   // Past due date
}

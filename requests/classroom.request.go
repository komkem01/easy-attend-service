package requests

import "github.com/google/uuid"

// CreateClassroomRequest for creating new classroom
type CreateClassroomRequest struct {
	SchoolID    *uuid.UUID `json:"school_id" validate:"omitempty,uuid"`
	Name        string     `json:"name" validate:"required,min=2,max=100"`
	Subject     string     `json:"subject" validate:"required,min=2,max=50"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	GradeLevel  *string    `json:"grade_level" validate:"omitempty,max=20"`
	Section     *string    `json:"section" validate:"omitempty,max=10"`
	RoomNumber  *string    `json:"room_number" validate:"omitempty,max=20"`
	MaxStudents *int       `json:"max_students" validate:"omitempty,min=1,max=200"`
	Schedule    *string    `json:"schedule" validate:"omitempty"`
}

// UpdateClassroomRequest for updating classroom
type UpdateClassroomRequest struct {
	SchoolID    *uuid.UUID `json:"school_id" validate:"omitempty,uuid"`
	Name        *string    `json:"name" validate:"omitempty,min=2,max=100"`
	Subject     *string    `json:"subject" validate:"omitempty,min=2,max=50"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	GradeLevel  *string    `json:"grade_level" validate:"omitempty,max=20"`
	Section     *string    `json:"section" validate:"omitempty,max=10"`
	RoomNumber  *string    `json:"room_number" validate:"omitempty,max=20"`
	MaxStudents *int       `json:"max_students" validate:"omitempty,min=1,max=200"`
	Schedule    *string    `json:"schedule" validate:"omitempty"`
	IsActive    *bool      `json:"is_active" validate:"omitempty"`
}

// ClassroomQueryRequest for filtering classrooms
type ClassroomQueryRequest struct {
	Page       int        `json:"page" query:"page" validate:"omitempty,min=1"`
	Limit      int        `json:"limit" query:"limit" validate:"omitempty,min=1,max=100"`
	Search     *string    `json:"search" query:"search" validate:"omitempty,min=1,max=100"`
	SchoolID   *uuid.UUID `json:"school_id" query:"school_id" validate:"omitempty,uuid"`
	TeacherID  *uuid.UUID `json:"teacher_id" query:"teacher_id" validate:"omitempty,uuid"`
	Subject    *string    `json:"subject" query:"subject" validate:"omitempty,max=50"`
	GradeLevel *string    `json:"grade_level" query:"grade_level" validate:"omitempty,max=20"`
	IsActive   *bool      `json:"is_active" query:"is_active" validate:"omitempty"`
}

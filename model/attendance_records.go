package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AttendanceRecords table structure
type AttendanceRecords struct {
	bun.BaseModel `bun:"table:attendance_records,alias:ar"`

	ID              uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SessionID       uuid.UUID  `json:"session_id" bun:"session_id,notnull,type:uuid"`
	StudentID       uuid.UUID  `json:"student_id" bun:"student_id,notnull,type:uuid"`
	Status          string     `json:"status" bun:"status,notnull,type:attendance_status"`
	CheckInTime     *time.Time `json:"check_in_time" bun:"check_in_time"`
	CheckInMethod   *string    `json:"check_in_method" bun:"check_in_method,type:check_in_method"`
	CheckInLocation *string    `json:"check_in_location" bun:"check_in_location"`
	LateMinutes     int        `json:"late_minutes" bun:"late_minutes,default:0"`
	Notes           *string    `json:"notes" bun:"notes"`
	MarkedBy        *uuid.UUID `json:"marked_by" bun:"marked_by,type:uuid"`
	IsModified      bool       `json:"is_modified" bun:"is_modified,notnull,default:false"`
	ModifiedAt      *time.Time `json:"modified_at" bun:"modified_at"`
	ModifiedBy      *uuid.UUID `json:"modified_by" bun:"modified_by,type:uuid"`
	CreatedAt       time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt       time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	Session  *AttendanceSessions `json:"session,omitempty" bun:"rel:belongs-to,join:session_id=id"`
	Student  *Users              `json:"student,omitempty" bun:"rel:belongs-to,join:student_id=id"`
	Marker   *Users              `json:"marker,omitempty" bun:"rel:belongs-to,join:marked_by=id"`
	Modifier *Users              `json:"modifier,omitempty" bun:"rel:belongs-to,join:modified_by=id"`
}

// TableName returns the table name
func (ar *AttendanceRecords) TableName() string {
	return "attendance_records"
}

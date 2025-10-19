package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AttendanceSessionsArchive table structure
type AttendanceSessionsArchive struct {
	bun.BaseModel `bun:"table:attendance_sessions_archive,alias:asa"`

	ID                   uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID          uuid.UUID  `json:"classroom_id" bun:"classroom_id,notnull,type:uuid"`
	Title                string     `json:"title" bun:"title,notnull"`
	Description          *string    `json:"description" bun:"description"`
	SessionDate          time.Time  `json:"session_date" bun:"session_date,notnull,type:date"`
	StartTime            time.Time  `json:"start_time" bun:"start_time,notnull,type:time"`
	EndTime              time.Time  `json:"end_time" bun:"end_time,notnull,type:time"`
	ActualStartTime      *time.Time `json:"actual_start_time" bun:"actual_start_time"`
	ActualEndTime        *time.Time `json:"actual_end_time" bun:"actual_end_time"`
	Status               string     `json:"status" bun:"status,notnull,default:'scheduled',type:session_status"`
	Method               string     `json:"method" bun:"method,notnull,default:'code',type:session_method"`
	SessionCode          *string    `json:"session_code" bun:"session_code"`
	QRCodeData           *string    `json:"qr_code_data" bun:"qr_code_data"`
	AllowLateCheck       bool       `json:"allow_late_check" bun:"allow_late_check,notnull,default:true"`
	LateThresholdMinutes int        `json:"late_threshold_minutes" bun:"late_threshold_minutes,default:15"`
	Location             *string    `json:"location" bun:"location"`
	Notes                *string    `json:"notes" bun:"notes"`
	CreatedBy            uuid.UUID  `json:"created_by" bun:"created_by,notnull,type:uuid"`
	CreatedAt            time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt            time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
}

// TableName returns the table name
func (asa *AttendanceSessionsArchive) TableName() string {
	return "attendance_sessions_archive"
}

package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AttendanceSessions table structure
type AttendanceSessions struct {
	bun.BaseModel `bun:"table:attendance_sessions,alias:as"`

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
	DeletedAt            *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	Classroom *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
	Creator   *Users      `json:"creator,omitempty" bun:"rel:belongs-to,join:created_by=id"`
}

// TableName returns the table name
func (as *AttendanceSessions) TableName() string {
	return "attendance_sessions"
}

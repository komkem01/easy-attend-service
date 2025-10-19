package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AttendanceAnalytics table structure
type AttendanceAnalytics struct {
	bun.BaseModel `bun:"table:attendance_analytics,alias:aa"`

	ID                 uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID        uuid.UUID `json:"classroom_id" bun:"classroom_id,notnull,type:uuid"`
	StudentID          uuid.UUID `json:"student_id" bun:"student_id,notnull,type:uuid"`
	MonthYear          string    `json:"month_year" bun:"month_year,notnull"`
	TotalSessions      int       `json:"total_sessions" bun:"total_sessions,default:0"`
	PresentCount       int       `json:"present_count" bun:"present_count,default:0"`
	AbsentCount        int       `json:"absent_count" bun:"absent_count,default:0"`
	LateCount          int       `json:"late_count" bun:"late_count,default:0"`
	ExcusedCount       int       `json:"excused_count" bun:"excused_count,default:0"`
	AttendanceRate     float64   `json:"attendance_rate" bun:"attendance_rate,default:0"`
	AverageLateMinutes float64   `json:"average_late_minutes" bun:"average_late_minutes,default:0"`
	CreatedAt          time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt          time.Time `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relations
	Classroom *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
	Student   *Users      `json:"student,omitempty" bun:"rel:belongs-to,join:student_id=id"`
}

// TableName returns the table name
func (aa *AttendanceAnalytics) TableName() string {
	return "attendance_analytics"
}

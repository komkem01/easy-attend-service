package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AcademicCalendar table structure
type AcademicCalendar struct {
	bun.BaseModel `bun:"table:academic_calendar,alias:ac"`

	ID                uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SchoolID          *uuid.UUID `json:"school_id" bun:"school_id,type:uuid"`
	Title             string     `json:"title" bun:"title,notnull"`
	Description       *string    `json:"description" bun:"description"`
	EventType         string     `json:"event_type" bun:"event_type,notnull,type:event_type"`
	StartDate         time.Time  `json:"start_date" bun:"start_date,notnull,type:date"`
	EndDate           *time.Time `json:"end_date" bun:"end_date,type:date"`
	IsRecurring       bool       `json:"is_recurring" bun:"is_recurring,notnull,default:false"`
	RecurrencePattern *string    `json:"recurrence_pattern" bun:"recurrence_pattern,type:jsonb"`
	AffectsAttendance bool       `json:"affects_attendance" bun:"affects_attendance,notnull,default:true"`
	CreatedBy         uuid.UUID  `json:"created_by" bun:"created_by,notnull,type:uuid"`
	CreatedAt         time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt         time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	School  *Schools `json:"school,omitempty" bun:"rel:belongs-to,join:school_id=id"`
	Creator *Users   `json:"creator,omitempty" bun:"rel:belongs-to,join:created_by=id"`
}

// TableName returns the table name
func (ac *AcademicCalendar) TableName() string {
	return "academic_calendar"
}

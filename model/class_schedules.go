package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ClassSchedules table structure
type ClassSchedules struct {
	bun.BaseModel `bun:"table:class_schedules,alias:csch"`

	ID             uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID    uuid.UUID  `json:"classroom_id" bun:"classroom_id,notnull,type:uuid"`
	DayOfWeek      int16      `json:"day_of_week" bun:"day_of_week,notnull"`
	StartTime      time.Time  `json:"start_time" bun:"start_time,notnull,type:time"`
	EndTime        time.Time  `json:"end_time" bun:"end_time,notnull,type:time"`
	RoomNumber     *string    `json:"room_number" bun:"room_number"`
	IsActive       bool       `json:"is_active" bun:"is_active,notnull,default:true"`
	EffectiveFrom  *time.Time `json:"effective_from" bun:"effective_from,type:date"`
	EffectiveUntil *time.Time `json:"effective_until" bun:"effective_until,type:date"`
	CreatedAt      time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`

	// Relations
	Classroom *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
}

// TableName returns the table name
func (cs *ClassSchedules) TableName() string {
	return "class_schedules"
}

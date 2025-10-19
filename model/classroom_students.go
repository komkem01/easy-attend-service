package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ClassroomStudents table structure
type ClassroomStudents struct {
	bun.BaseModel `bun:"table:classroom_students,alias:cs"`

	ID            uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID   uuid.UUID `json:"classroom_id" bun:"classroom_id,notnull,type:uuid"`
	StudentID     uuid.UUID `json:"student_id" bun:"student_id,notnull,type:uuid"`
	StudentNumber *string   `json:"student_number" bun:"student_number"`
	SeatNumber    *string   `json:"seat_number" bun:"seat_number"`
	EnrolledAt    time.Time `json:"enrolled_at" bun:"enrolled_at,notnull,default:now()"`
	IsActive      bool      `json:"is_active" bun:"is_active,notnull,default:true"`
	Notes         *string   `json:"notes" bun:"notes"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`

	// Relations
	Classroom *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
	Student   *Users      `json:"student,omitempty" bun:"rel:belongs-to,join:student_id=id"`
}

// TableName returns the table name
func (cs *ClassroomStudents) TableName() string {
	return "classroom_students"
}

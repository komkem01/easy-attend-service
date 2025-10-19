package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Classrooms table structure
type Classrooms struct {
	bun.BaseModel `bun:"table:classrooms,alias:c"`

	ID            uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SchoolID      *uuid.UUID `json:"school_id" bun:"school_id,type:uuid"`
	Name          string     `json:"name" bun:"name,notnull"`
	Subject       string     `json:"subject" bun:"subject,notnull"`
	Description   *string    `json:"description" bun:"description"`
	GradeLevel    *string    `json:"grade_level" bun:"grade_level"`
	Section       *string    `json:"section" bun:"section"`
	RoomNumber    *string    `json:"room_number" bun:"room_number"`
	TeacherID     uuid.UUID  `json:"teacher_id" bun:"teacher_id,notnull,type:uuid"`
	ClassroomCode string     `json:"classroom_code" bun:"classroom_code,notnull,unique"`
	MaxStudents   int        `json:"max_students" bun:"max_students,default:50"`
	Schedule      *string    `json:"schedule" bun:"schedule,type:jsonb"`
	IsActive      bool       `json:"is_active" bun:"is_active,notnull,default:true"`
	CreatedAt     time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt     time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	School             *Schools              `json:"school,omitempty" bun:"rel:belongs-to,join:school_id=id"`
	Teacher            *Users                `json:"teacher,omitempty" bun:"rel:belongs-to,join:teacher_id=id"`
	ClassroomStudents  []*ClassroomStudents  `json:"classroom_students,omitempty" bun:"rel:has-many,join:id=classroom_id"`
	AttendanceSessions []*AttendanceSessions `json:"attendance_sessions,omitempty" bun:"rel:has-many,join:id=classroom_id"`
	ClassSchedules     []*ClassSchedules     `json:"class_schedules,omitempty" bun:"rel:has-many,join:id=classroom_id"`
	Assignments        []*Assignments        `json:"assignments,omitempty" bun:"rel:has-many,join:id=classroom_id"`
	Messages           []*Messages           `json:"messages,omitempty" bun:"rel:has-many,join:id=classroom_id"`
}

// TableName returns the table name
func (c *Classrooms) TableName() string {
	return "classrooms"
}

package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ClassroomMembers table structure
type ClassroomMembers struct {
	bun.BaseModel `bun:"table:classroom_members,alias:cm"`

	ID          uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID uuid.UUID  `json:"classroom_id" bun:"classroom_id,notnull,type:uuid"`
	UserID      uuid.UUID  `json:"user_id" bun:"user_id,notnull,type:uuid"`
	Role        string     `json:"role" bun:"role,notnull,default:'student',type:classroom_role"`
	Status      string     `json:"status" bun:"status,notnull,default:'active',type:member_status"`
	JoinedAt    time.Time  `json:"joined_at" bun:"joined_at,notnull,default:now()"`
	LeftAt      *time.Time `json:"left_at" bun:"left_at"`
	CreatedAt   time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relationships
	Classroom *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
	User      *Users      `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (cm *ClassroomMembers) TableName() string {
	return "classroom_members"
}

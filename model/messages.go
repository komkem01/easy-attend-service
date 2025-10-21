package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Messages table structure
type Messages struct {
	bun.BaseModel `bun:"table:messages,alias:m"`

	ID              uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SenderID        uuid.UUID  `json:"sender_id" bun:"sender_id,notnull,type:uuid"`
	RecipientID     *uuid.UUID `json:"recipient_id" bun:"recipient_id,type:uuid"`
	ClassroomID     *uuid.UUID `json:"classroom_id" bun:"classroom_id,type:uuid"`
	ParentMessageID *uuid.UUID `json:"parent_message_id" bun:"parent_message_id,type:uuid"`
	Subject         *string    `json:"subject" bun:"subject"`
	Content         string     `json:"content" bun:"content,notnull"`
	MessageType     string     `json:"message_type" bun:"message_type,notnull,default:'private',type:message_type"`
	Priority        string     `json:"priority" bun:"priority,notnull,default:'normal',type:priority_level"`
	IsRead          bool       `json:"is_read" bun:"is_read,notnull,default:false"`
	ReadAt          *time.Time `json:"read_at" bun:"read_at"`
	IsDeleted       bool       `json:"is_deleted" bun:"is_deleted,notnull,default:false"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`
	CreatedAt       time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`

	// Relations
	Sender        *Users      `json:"sender,omitempty" bun:"rel:belongs-to,join:sender_id=id"`
	Recipient     *Users      `json:"recipient,omitempty" bun:"rel:belongs-to,join:recipient_id=id"`
	Classroom     *Classrooms `json:"classroom,omitempty" bun:"rel:belongs-to,join:classroom_id=id"`
	ParentMessage *Messages   `json:"parent_message,omitempty" bun:"rel:belongs-to,join:parent_message_id=id"`
}

// TableName returns the table name
func (m *Messages) TableName() string {
	return "messages"
}

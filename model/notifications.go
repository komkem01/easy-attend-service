package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Notifications table structure
type Notifications struct {
	bun.BaseModel `bun:"table:notifications,alias:n"`

	ID                uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID            uuid.UUID  `json:"user_id" bun:"user_id,notnull,type:uuid"`
	Type              string     `json:"type" bun:"type,notnull,type:notification_type"`
	Title             string     `json:"title" bun:"title,notnull"`
	Message           string     `json:"message" bun:"message,notnull"`
	Data              *string    `json:"data" bun:"data,type:jsonb"`
	IsRead            bool       `json:"is_read" bun:"is_read,notnull,default:false"`
	ReadAt            *time.Time `json:"read_at" bun:"read_at"`
	ReferenceType     *string    `json:"reference_type" bun:"reference_type,type:reference_type"`
	ReferenceID       *uuid.UUID `json:"reference_id" bun:"reference_id,type:uuid"`
	ScheduledFor      *time.Time `json:"scheduled_for" bun:"scheduled_for"`
	SentAt            *time.Time `json:"sent_at" bun:"sent_at"`
	DeliveryStatus    string     `json:"delivery_status" bun:"delivery_status,notnull,default:'pending',type:delivery_status"`
	DeliveryChannel   string     `json:"delivery_channel" bun:"delivery_channel,notnull,default:'in_app',type:delivery_channel"`
	ExternalMessageID *string    `json:"external_message_id" bun:"external_message_id"`
	FailureReason     *string    `json:"failure_reason" bun:"failure_reason"`
	RetryCount        int        `json:"retry_count" bun:"retry_count,default:0"`
	LastRetryAt       *time.Time `json:"last_retry_at" bun:"last_retry_at"`
	ExpiresAt         *time.Time `json:"expires_at" bun:"expires_at"`
	CreatedAt         time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt         time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relationships
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (n *Notifications) TableName() string {
	return "notifications"
}

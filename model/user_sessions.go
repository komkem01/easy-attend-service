package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// UserSessions table structure
type UserSessions struct {
	bun.BaseModel `bun:"table:user_sessions,alias:us"`

	ID           uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID       uuid.UUID `json:"user_id" bun:"user_id,notnull,type:uuid"`
	SessionID    string    `json:"session_id" bun:"session_id,notnull,unique"`
	IPAddress    *string   `json:"ip_address" bun:"ip_address"`
	UserAgent    *string   `json:"user_agent" bun:"user_agent"`
	DeviceInfo   *string   `json:"device_info" bun:"device_info,type:jsonb"`
	Location     *string   `json:"location" bun:"location,type:jsonb"`
	IsActive     bool      `json:"is_active" bun:"is_active,notnull,default:true"`
	LastActivity time.Time `json:"last_activity" bun:"last_activity,notnull,default:now()"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`
	ExpiresAt    time.Time `json:"expires_at" bun:"expires_at,notnull"`

	// Relationships
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (us *UserSessions) TableName() string {
	return "user_sessions"
}

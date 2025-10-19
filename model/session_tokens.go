package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// SessionTokens table structure
type SessionTokens struct {
	bun.BaseModel `bun:"table:session_tokens,alias:st"`

	ID           uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID       uuid.UUID  `json:"user_id" bun:"user_id,notnull,type:uuid"`
	TokenHash    string     `json:"token_hash" bun:"token_hash,notnull,unique"`
	DeviceInfo   *string    `json:"device_info" bun:"device_info,type:jsonb"`
	IPAddress    *string    `json:"ip_address" bun:"ip_address"`
	UserAgent    *string    `json:"user_agent" bun:"user_agent"`
	LastUsedAt   time.Time  `json:"last_used_at" bun:"last_used_at,notnull,default:now()"`
	ExpiresAt    time.Time  `json:"expires_at" bun:"expires_at,notnull"`
	IsRevoked    bool       `json:"is_revoked" bun:"is_revoked,notnull,default:false"`
	RevokedAt    *time.Time `json:"revoked_at" bun:"revoked_at"`
	RevokeReason *string    `json:"revoke_reason" bun:"revoke_reason"`
	CreatedAt    time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt    time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relationships
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (st *SessionTokens) TableName() string {
	return "session_tokens"
}

package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// SecurityEvents table structure
type SecurityEvents struct {
	bun.BaseModel `bun:"table:security_events,alias:se"`

	ID          uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID      *uuid.UUID `json:"user_id" bun:"user_id,type:uuid"`
	EventType   string     `json:"event_type" bun:"event_type,notnull"`
	Description string     `json:"description" bun:"description,notnull"`
	RiskLevel   string     `json:"risk_level" bun:"risk_level,notnull,default:'low',type:risk_level"`
	IPAddress   *string    `json:"ip_address" bun:"ip_address"`
	UserAgent   *string    `json:"user_agent" bun:"user_agent"`
	Location    *string    `json:"location" bun:"location,type:jsonb"`
	Details     *string    `json:"details" bun:"details,type:jsonb"`
	IsResolved  bool       `json:"is_resolved" bun:"is_resolved,notnull,default:false"`
	ResolvedAt  *time.Time `json:"resolved_at" bun:"resolved_at"`
	ResolvedBy  *uuid.UUID `json:"resolved_by" bun:"resolved_by,type:uuid"`
	ActionTaken *string    `json:"action_taken" bun:"action_taken"`
	CreatedAt   time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relationships
	User     *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
	Resolver *Users `json:"resolver,omitempty" bun:"rel:belongs-to,join:resolved_by=id"`
}

// TableName returns the table name
func (se *SecurityEvents) TableName() string {
	return "security_events"
}

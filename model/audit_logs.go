package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AuditLogs table structure
type AuditLogs struct {
	bun.BaseModel `bun:"table:audit_logs,alias:al"`

	ID             uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID         *uuid.UUID `json:"user_id" bun:"user_id,type:uuid"`
	Action         string     `json:"action" bun:"action,notnull"`
	Table          string     `json:"table_name" bun:"table_name,notnull"`
	RecordID       *uuid.UUID `json:"record_id" bun:"record_id,type:uuid"`
	OldValues      *string    `json:"old_values" bun:"old_values,type:jsonb"`
	NewValues      *string    `json:"new_values" bun:"new_values,type:jsonb"`
	IPAddress      *string    `json:"ip_address" bun:"ip_address"`
	UserAgent      *string    `json:"user_agent" bun:"user_agent"`
	AdditionalData *string    `json:"additional_data" bun:"additional_data,type:jsonb"`
	CreatedAt      time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`

	// Relationships
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (al *AuditLogs) TableName() string {
	return "audit_logs"
}

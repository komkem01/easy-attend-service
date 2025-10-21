package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// SearchLogs table structure
type SearchLogs struct {
	bun.BaseModel `bun:"table:search_logs,alias:sl"`

	ID           uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID       *uuid.UUID `json:"user_id" bun:"user_id,type:uuid"`
	SearchType   string     `json:"search_type" bun:"search_type,notnull,type:search_type"`
	Query        string     `json:"query" bun:"query,notnull"`
	Filters      *string    `json:"filters" bun:"filters,type:jsonb"`
	ResultCount  int        `json:"result_count" bun:"result_count,default:0"`
	ResponseTime float64    `json:"response_time" bun:"response_time"`
	IPAddress    *string    `json:"ip_address" bun:"ip_address"`
	UserAgent    *string    `json:"user_agent" bun:"user_agent"`
	CreatedAt    time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relationships
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (sl *SearchLogs) TableName() string {
	return "search_logs"
}

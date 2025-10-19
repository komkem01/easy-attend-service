package model

import (
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ApiRateLimits table structure
type ApiRateLimits struct {
	bun.BaseModel `bun:"table:api_rate_limits,alias:arl"`

	ID           uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID       *uuid.UUID `json:"user_id" bun:"user_id,type:uuid"`
	IpAddress    *net.IP    `json:"ip_address" bun:"ip_address,type:inet"`
	Endpoint     string     `json:"endpoint" bun:"endpoint,notnull"`
	RequestCount int        `json:"request_count" bun:"request_count,default:1"`
	WindowStart  time.Time  `json:"window_start" bun:"window_start,notnull,default:now()"`
	IsBlocked    bool       `json:"is_blocked" bun:"is_blocked,notnull,default:false"`
	BlockedUntil *time.Time `json:"blocked_until" bun:"blocked_until"`

	// Relations
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (arl *ApiRateLimits) TableName() string {
	return "api_rate_limits"
}

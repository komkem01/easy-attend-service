package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ApiKeys table structure
type ApiKeys struct {
	bun.BaseModel `bun:"table:api_keys,alias:ak"`

	ID          uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID      uuid.UUID  `json:"user_id" bun:"user_id,notnull,type:uuid"`
	KeyName     string     `json:"key_name" bun:"key_name,notnull"`
	ApiKey      string     `json:"api_key" bun:"api_key,notnull,unique"`
	ApiSecret   string     `json:"-" bun:"api_secret,notnull"`
	Permissions *string    `json:"permissions" bun:"permissions,type:jsonb"`
	IsActive    bool       `json:"is_active" bun:"is_active,notnull,default:true"`
	ExpiresAt   *time.Time `json:"expires_at" bun:"expires_at"`
	LastUsedAt  *time.Time `json:"last_used_at" bun:"last_used_at"`
	UsageCount  int        `json:"usage_count" bun:"usage_count,default:0"`
	CreatedAt   time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (ak *ApiKeys) TableName() string {
	return "api_keys"
}

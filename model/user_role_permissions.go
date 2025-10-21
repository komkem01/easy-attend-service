package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// UserRolePermissions table structure
type UserRolePermissions struct {
	bun.BaseModel `bun:"table:user_role_permissions,alias:urp"`

	ID         uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID     uuid.UUID  `json:"user_id" bun:"user_id,notnull,type:uuid"`
	Permission string     `json:"permission" bun:"permission,notnull,type:permission_type"`
	Resource   *string    `json:"resource" bun:"resource"`
	Context    *string    `json:"context" bun:"context,type:jsonb"`
	GrantedBy  uuid.UUID  `json:"granted_by" bun:"granted_by,notnull,type:uuid"`
	GrantedAt  time.Time  `json:"granted_at" bun:"granted_at,notnull,default:now()"`
	ExpiresAt  *time.Time `json:"expires_at" bun:"expires_at"`
	IsActive   bool       `json:"is_active" bun:"is_active,notnull,default:true"`
	CreatedAt  time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt  time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relationships
	User    *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
	Granter *Users `json:"granter,omitempty" bun:"rel:belongs-to,join:granted_by=id"`
}

// TableName returns the table name
func (urp *UserRolePermissions) TableName() string {
	return "user_role_permissions"
}

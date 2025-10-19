package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Schools table structure
type Schools struct {
	bun.BaseModel `bun:"table:schools,alias:s"`

	ID         uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name       string    `json:"name" bun:"name,notnull"`
	Address    *string   `json:"address" bun:"address"`
	Phone      *string   `json:"phone" bun:"phone"`
	Email      *string   `json:"email" bun:"email"`
	WebsiteURL *string   `json:"website_url" bun:"website_url"`
	LogoURL    *string   `json:"logo_url" bun:"logo_url"`
	IsActive   bool      `json:"is_active" bun:"is_active,notnull,default:true"`
	CreatedAt  time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt  time.Time `json:"updated_at" bun:"updated_at,notnull,default:now()"`

	// Relations
	Users      []*Users      `json:"users,omitempty" bun:"rel:has-many,join:id=school_id"`
	Classrooms []*Classrooms `json:"classrooms,omitempty" bun:"rel:has-many,join:id=school_id"`
}

// TableName returns the table name
func (s *Schools) TableName() string {
	return "schools"
}

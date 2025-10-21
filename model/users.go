package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Users table structure
type Users struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID            uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SchoolID      *uuid.UUID `json:"school_id" bun:"school_id,type:uuid"`
	Username      string     `json:"username" bun:"username,notnull,unique"`
	Email         string     `json:"email" bun:"email,notnull,unique"`
	PasswordHash  string     `json:"-" bun:"password_hash,notnull"`
	PrefixID      *int       `json:"prefix_id" bun:"prefix_id"`
	FirstName     string     `json:"first_name" bun:"first_name,notnull"`
	LastName      string     `json:"last_name" bun:"last_name,notnull"`
	GenderID      *int       `json:"gender_id" bun:"gender_id"`
	Role          string     `json:"role" bun:"role,notnull,type:user_role"`
	AvatarURL     *string    `json:"avatar_url" bun:"avatar_url"`
	Phone         *string    `json:"phone" bun:"phone"`
	DateOfBirth   *time.Time `json:"date_of_birth" bun:"date_of_birth,type:date"`
	Address       *string    `json:"address" bun:"address"`
	IsActive      bool       `json:"is_active" bun:"is_active,notnull,default:true"`
	EmailVerified bool       `json:"email_verified" bun:"email_verified,notnull,default:false"`
	LastLoginAt   *time.Time `json:"last_login_at" bun:"last_login_at"`
	CreatedAt     time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt     time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	School *Schools  `json:"school,omitempty" bun:"rel:belongs-to,join:school_id=id"`
	Prefix *Prefixes `json:"prefix,omitempty" bun:"rel:belongs-to,join:prefix_id=id"`
	Gender *Genders  `json:"gender,omitempty" bun:"rel:belongs-to,join:gender_id=id"`
}

// TableName returns the table name
func (u *Users) TableName() string {
	return "users"
}

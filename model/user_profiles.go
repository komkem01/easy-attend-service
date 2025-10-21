package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// UserProfiles table structure
type UserProfiles struct {
	bun.BaseModel `bun:"table:user_profiles,alias:up"`

	ID               uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID           uuid.UUID  `json:"user_id" bun:"user_id,notnull,unique,type:uuid"`
	FirstName        *string    `json:"first_name" bun:"first_name"`
	LastName         *string    `json:"last_name" bun:"last_name"`
	FullName         *string    `json:"full_name" bun:"full_name"`
	DateOfBirth      *time.Time `json:"date_of_birth" bun:"date_of_birth,type:date"`
	Gender           *string    `json:"gender" bun:"gender,type:gender"`
	PhoneNumber      *string    `json:"phone_number" bun:"phone_number"`
	Address          *string    `json:"address" bun:"address"`
	City             *string    `json:"city" bun:"city"`
	State            *string    `json:"state" bun:"state"`
	PostalCode       *string    `json:"postal_code" bun:"postal_code"`
	Country          *string    `json:"country" bun:"country"`
	ProfilePicture   *string    `json:"profile_picture" bun:"profile_picture"`
	Bio              *string    `json:"bio" bun:"bio"`
	Website          *string    `json:"website" bun:"website"`
	SocialLinks      *string    `json:"social_links" bun:"social_links,type:jsonb"`
	Preferences      *string    `json:"preferences" bun:"preferences,type:jsonb"`
	EmergencyContact *string    `json:"emergency_contact" bun:"emergency_contact,type:jsonb"`
	CreatedAt        time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt        time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relationships
	User *Users `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}

// TableName returns the table name
func (up *UserProfiles) TableName() string {
	return "user_profiles"
}

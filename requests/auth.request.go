package requests

// LoginRequest represents the login request structure
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest represents the registration request structure
type RegisterRequest struct {
	UserType        string  `json:"user_type" binding:"required,oneof=student teacher admin"`
	Prefix          *string `json:"prefix" binding:"omitempty"`
	FirstName       string  `json:"first_name" binding:"required,min=1,max=50"`
	LastName        string  `json:"last_name" binding:"required,min=1,max=50"`
	Gender          *string `json:"gender" binding:"omitempty,oneof=ชาย หญิง อื่นๆ Male Female Other"`
	Username        string  `json:"username" binding:"required,min=3,max=50"`
	SchoolName      string  `json:"school_name" binding:"required,min=2,max=200"`
	Email           string  `json:"email" binding:"required,email"`
	Password        string  `json:"password" binding:"required,min=6"`
	ConfirmPassword string  `json:"confirm_password" binding:"required,min=6"`
	Phone           *string `json:"phone" binding:"omitempty"`
}

// UpdateProfileRequest represents the profile update request structure (for PATCH)
type UpdateProfileRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	FullName string `json:"full_name" binding:"omitempty,min=2,max=100"`
}

// ReplaceProfileRequest represents the full profile replacement request structure (for PUT)
type ReplaceProfileRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	FullName string `json:"full_name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
}

// RefreshTokenRequest represents the refresh token request structure
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserProfileRequest represents user profile update request
type UserProfileRequest struct {
	Prefix         *string `json:"prefix" binding:"omitempty"`
	FirstName      *string `json:"first_name" binding:"omitempty,max=50"`
	LastName       *string `json:"last_name" binding:"omitempty,max=50"`
	Gender         *string `json:"gender" binding:"omitempty,oneof=ชาย หญิง อื่นๆ Male Female Other"`
	FullName       *string `json:"full_name" binding:"omitempty,max=100"`
	PhoneNumber    *string `json:"phone_number" binding:"omitempty"`
	DateOfBirth    *string `json:"date_of_birth" binding:"omitempty"` // Format: YYYY-MM-DD
	Address        *string `json:"address" binding:"omitempty,max=255"`
	City           *string `json:"city" binding:"omitempty,max=100"`
	State          *string `json:"state" binding:"omitempty,max=100"`
	PostalCode     *string `json:"postal_code" binding:"omitempty,max=20"`
	Country        *string `json:"country" binding:"omitempty,max=100"`
	Bio            *string `json:"bio" binding:"omitempty,max=500"`
	Website        *string `json:"website" binding:"omitempty,url"`
	ProfilePicture *string `json:"profile_picture" binding:"omitempty,url"`
}

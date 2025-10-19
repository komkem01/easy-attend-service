package requests

// LoginRequest represents the login request structure
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest represents the registration request structure
type RegisterRequest struct {
	Email      string  `json:"email" binding:"required,email"`
	Password   string  `json:"password" binding:"required,min=6"`
	Name       string  `json:"name" binding:"required,min=2,max=100"`
	Role       string  `json:"role" binding:"required,oneof=student teacher admin"`
	SchoolName string  `json:"school_name" binding:"required,min=2,max=200"`
	Phone      *string `json:"phone" binding:"omitempty"`
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

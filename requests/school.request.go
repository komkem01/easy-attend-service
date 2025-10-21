package requests

// CreateSchoolRequest represents the school creation request structure
type CreateSchoolRequest struct {
	Name       string  `json:"name" binding:"required,min=2,max=200"`
	Address    *string `json:"address" binding:"omitempty,max=500"`
	Phone      *string `json:"phone" binding:"omitempty"`
	Email      *string `json:"email" binding:"omitempty,email"`
	WebsiteURL *string `json:"website_url" binding:"omitempty,url"`
}

// UpdateSchoolRequest represents the school update request structure
type UpdateSchoolRequest struct {
	Name       *string `json:"name" binding:"omitempty,min=2,max=200"`
	Address    *string `json:"address" binding:"omitempty,max=500"`
	Phone      *string `json:"phone" binding:"omitempty"`
	Email      *string `json:"email" binding:"omitempty,email"`
	WebsiteURL *string `json:"website_url" binding:"omitempty,url"`
}

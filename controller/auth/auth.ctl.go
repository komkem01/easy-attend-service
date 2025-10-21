package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/komkem01/easy-attend-service/model"
	"github.com/komkem01/easy-attend-service/requests"
	"github.com/komkem01/easy-attend-service/response"
)

// Login handles user login
func Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Call login service
	user, err := LoginUserService(c.Request.Context(), req)
	if err != nil {
		response.Unauthorized(c, "Invalid email or password")
		return
	}

	// Generate auth response with tokens
	authResponse, err := AuthResponseService(user)
	if err != nil {
		response.InternalServerError(c, "Failed to generate authentication tokens")
		return
	}

	response.Success(c, authResponse)
}

// Register handles user registration
func Register(c *gin.Context) {
	var req requests.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Call register service
	user, err := RegisterUserService(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "email already registered" {
			response.Conflict(c, "Email already registered")
			return
		}
		response.InternalServerError(c, "Failed to create user account")
		return
	}

	// Generate auth response with tokens
	authResponse, err := AuthResponseService(user)
	if err != nil {
		response.InternalServerError(c, "Failed to generate authentication tokens")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
		"data":    authResponse,
	})
}

// RefreshToken handles token refresh
func RefreshToken(c *gin.Context) {
	var req requests.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Call refresh token service
	tokenResponse, err := RefreshTokenService(c.Request.Context(), req.RefreshToken)
	if err != nil {
		response.Unauthorized(c, "Invalid or expired refresh token")
		return
	}

	response.Success(c, tokenResponse)
}

// GetProfile handles getting user profile
func GetProfile(c *gin.Context) {
	// Get user ID from JWT token
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	// Get user and profile information
	user, profile, err := GetUserProfileService(c.Request.Context(), userID.(string))
	if err != nil {
		response.NotFound(c, "User profile not found")
		return
	}

	// Prepare response data
	userData := map[string]interface{}{
		"id":            user.ID,
		"username":      user.Username,
		"email":         user.Email,
		"first_name":    user.FirstName,
		"last_name":     user.LastName,
		"full_name":     user.FirstName + " " + user.LastName,
		"role":          user.Role,
		"is_active":     user.IsActive,
		"created_at":    user.CreatedAt,
		"phone":         user.Phone,
		"date_of_birth": user.DateOfBirth,
		"address":       user.Address,
	}

	// Add prefix information if available
	if user.Prefix != nil {
		userData["prefix"] = map[string]interface{}{
			"id":           user.Prefix.ID,
			"code":         user.Prefix.Code,
			"name_th":      user.Prefix.NameTH,
			"name_en":      user.Prefix.NameEN,
			"abbreviation": user.Prefix.Abbreviation,
		}
	}

	// Add gender information if available
	if user.Gender != nil {
		userData["gender"] = map[string]interface{}{
			"id":           user.Gender.ID,
			"code":         user.Gender.Code,
			"name_th":      user.Gender.NameTH,
			"name_en":      user.Gender.NameEN,
			"abbreviation": user.Gender.Abbreviation,
		}
	}

	// Add school information if available
	if user.School != nil {
		userData["school"] = map[string]interface{}{
			"id":   user.School.ID,
			"name": user.School.Name,
		}
	}

	// Add profile information if available
	if profile != nil {
		userData["profile"] = profile
	}

	response.Success(c, userData)
}

// UpdateProfile handles updating user profile (PATCH method for partial updates)
func UpdateProfile(c *gin.Context) {
	var req requests.UserProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Get user ID from JWT token
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	// Convert request to map for flexible updating
	profileData := make(map[string]interface{})
	userData := make(map[string]interface{})

	// User data updates
	if req.Prefix != nil {
		prefixID, err := FindPrefixIDByName(c.Request.Context(), *req.Prefix)
		if err == nil && prefixID != nil {
			userData["prefix_id"] = *prefixID
		}
	}
	if req.FirstName != nil {
		userData["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		userData["last_name"] = *req.LastName
	}
	if req.Gender != nil {
		genderID, err := FindGenderIDByName(c.Request.Context(), *req.Gender)
		if err == nil && genderID != nil {
			userData["gender_id"] = *genderID
		}
	}

	// Profile data updates
	if req.FullName != nil {
		profileData["full_name"] = *req.FullName
	}
	if req.PhoneNumber != nil {
		profileData["phone_number"] = *req.PhoneNumber
	}
	if req.DateOfBirth != nil {
		profileData["date_of_birth"] = *req.DateOfBirth
	}
	if req.Address != nil {
		profileData["address"] = *req.Address
	}
	if req.City != nil {
		profileData["city"] = *req.City
	}
	if req.State != nil {
		profileData["state"] = *req.State
	}
	if req.PostalCode != nil {
		profileData["postal_code"] = *req.PostalCode
	}
	if req.Country != nil {
		profileData["country"] = *req.Country
	}
	if req.Bio != nil {
		profileData["bio"] = *req.Bio
	}
	if req.Website != nil {
		profileData["website"] = *req.Website
	}
	if req.ProfilePicture != nil {
		profileData["profile_picture"] = *req.ProfilePicture
	}

	// Update user data if any user fields are provided
	if len(userData) > 0 {
		err := UpdateUserDataService(c.Request.Context(), userID.(string), userData)
		if err != nil {
			response.InternalServerError(c, "Failed to update user data")
			return
		}
	}

	// Update profile
	profile, err := UpdateUserProfileService(c.Request.Context(), userID.(string), profileData)
	if err != nil {
		response.InternalServerError(c, "Failed to update profile")
		return
	}

	response.Success(c, gin.H{
		"message": "Profile updated successfully",
		"profile": profile,
	})
}

// ReplaceProfile handles full profile replacement (PUT method for complete replacement)
func ReplaceProfile(c *gin.Context) {
	var req requests.ReplaceProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// TODO: Implement full profile replacement logic using PUT method
	response.Success(c, gin.H{
		"message": "Profile full replacement functionality to be implemented",
		"method":  "PUT",
	})
}

// GetGenders handles getting all genders
func GetGenders(c *gin.Context) {
	genders, err := GetGendersService(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "Failed to get genders")
		return
	}

	response.Success(c, genders)
}
func GetGendersByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Gender ID is required")
		return
	}

	genderID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(c, "Invalid gender ID format")
		return
	}

	gender, err := GetGendersServiceByID(c.Request.Context(), genderID)
	if err != nil {
		response.InternalServerError(c, "Failed to get gender")
		return
	}

	response.Success(c, gender)
}

// GetPrefixes handles getting all prefixes
func GetPrefixes(c *gin.Context) {
	// Check if filtering by gender
	genderCode := c.Query("gender")

	var prefixes []model.Prefixes
	var err error

	if genderCode != "" {
		prefixes, err = GetPrefixesByGenderService(c.Request.Context(), genderCode)
	} else {
		prefixes, err = GetPrefixesService(c.Request.Context())
	}

	if err != nil {
		response.InternalServerError(c, "Failed to get prefixes")
		return
	}

	response.Success(c, prefixes)
}

func GetPrefixesByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Prefix ID is required")
		return
	}

	prefixID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(c, "Invalid prefix ID format")
		return
	}

	prefix, err := GetPrefixesServiceByID(c.Request.Context(), prefixID)
	if err != nil {
		response.InternalServerError(c, "Failed to get prefix")
		return
	}

	response.Success(c, prefix)
}

// Schools CRUD Operations

// GetSchools handles getting all schools
func GetSchools(c *gin.Context) {
	// Check if searching
	searchQuery := c.Query("search")

	var schools []model.Schools
	var err error

	if searchQuery != "" {
		schools, err = SearchSchoolsService(c.Request.Context(), searchQuery)
	} else {
		schools, err = GetSchoolsService(c.Request.Context())
	}

	if err != nil {
		response.InternalServerError(c, "Failed to get schools")
		return
	}

	response.Success(c, schools)
}

// GetSchoolByID handles getting school by ID
func GetSchoolByID(c *gin.Context) {
	schoolID := c.Param("id")
	if schoolID == "" {
		response.BadRequest(c, "School ID is required")
		return
	}

	school, err := GetSchoolByIDService(c.Request.Context(), schoolID)
	if err != nil {
		response.NotFound(c, "School not found")
		return
	}

	response.Success(c, school)
}

// CreateSchool handles creating a new school
func CreateSchool(c *gin.Context) {
	var req requests.CreateSchoolRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Convert optional fields to empty strings if nil
	address := ""
	if req.Address != nil {
		address = *req.Address
	}
	phone := ""
	if req.Phone != nil {
		phone = *req.Phone
	}
	email := ""
	if req.Email != nil {
		email = *req.Email
	}
	websiteURL := ""
	if req.WebsiteURL != nil {
		websiteURL = *req.WebsiteURL
	}

	school, err := CreateSchoolService(c.Request.Context(), req.Name, address, phone, email, websiteURL)
	if err != nil {
		if err.Error() == "school name already exists" {
			response.Conflict(c, "School name already exists")
			return
		}
		response.InternalServerError(c, "Failed to create school")
		return
	}

	response.Created(c, school)
}

// UpdateSchool handles updating school information
func UpdateSchool(c *gin.Context) {
	schoolID := c.Param("id")
	if schoolID == "" {
		response.BadRequest(c, "School ID is required")
		return
	}

	var req requests.UpdateSchoolRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Convert request to map for flexible updating
	updateData := make(map[string]interface{})

	if req.Name != nil {
		updateData["name"] = *req.Name
	}
	if req.Address != nil {
		updateData["address"] = *req.Address
	}
	if req.Phone != nil {
		updateData["phone"] = *req.Phone
	}
	if req.Email != nil {
		updateData["email"] = *req.Email
	}
	if req.WebsiteURL != nil {
		updateData["website_url"] = *req.WebsiteURL
	}

	school, err := UpdateSchoolService(c.Request.Context(), schoolID, updateData)
	if err != nil {
		if err.Error() == "school name already exists" {
			response.Conflict(c, "School name already exists")
			return
		}
		if err.Error() == "no data to update" {
			response.BadRequest(c, "No data to update")
			return
		}
		response.InternalServerError(c, "Failed to update school")
		return
	}

	response.Success(c, school)
}

// DeleteSchool handles soft deleting a school
func DeleteSchool(c *gin.Context) {
	schoolID := c.Param("id")
	if schoolID == "" {
		response.BadRequest(c, "School ID is required")
		return
	}

	err := DeleteSchoolService(c.Request.Context(), schoolID)
	if err != nil {
		if err.Error() == "cannot delete school that has active users" {
			response.Conflict(c, "Cannot delete school that has active users")
			return
		}
		response.InternalServerError(c, "Failed to delete school")
		return
	}

	response.Success(c, gin.H{"message": "School deleted successfully"})
}

// GetInfo handles getting system information
func GetInfo(c *gin.Context) {
	info := map[string]interface{}{
		"service_name":  "Easy Attend Service",
		"version":       "1.0.0",
		"description":   "A comprehensive attendance management system for educational institutions",
		"api_version":   "v1",
		"documentation": "/docs",
		"health_check":  "/health",
		"features": []string{
			"User Authentication & Authorization",
			"School Management",
			"Student Registration",
			"Teacher Management",
			"Attendance Tracking",
			"Profile Management",
			"Gender & Prefix Support",
			"Multi-language Support (Thai/English)",
		},
		"endpoints": map[string]interface{}{
			"auth": map[string]string{
				"login":    "POST /api/v1/auth/login",
				"register": "POST /api/v1/auth/register",
				"refresh":  "POST /api/v1/auth/refresh",
			},
			"profile": map[string]string{
				"get_own":   "GET /api/v1/profile",
				"update":    "PATCH /api/v1/profile",
				"get_by_id": "GET /api/v1/profile/{id}",
			},
			"schools": map[string]string{
				"list":   "GET /api/v1/schools",
				"get":    "GET /api/v1/schools/{id}",
				"create": "POST /api/v1/schools",
				"update": "PATCH /api/v1/schools/{id}",
				"delete": "DELETE /api/v1/schools/{id}",
				"search": "GET /api/v1/schools?search={query}",
			},
			"reference": map[string]string{
				"genders":  "GET /api/v1/genders",
				"prefixes": "GET /api/v1/prefixes",
			},
		},
		"contact": map[string]string{
			"developer": "Easy Attend Development Team",
			"email":     "support@easyattend.com",
		},
	}

	response.Success(c, info)
}

// GetProfileByID handles getting user profile by user ID
func GetProfileByID(c *gin.Context) {
	// Get user ID from URL parameter
	userID := c.Param("id")
	if userID == "" {
		response.BadRequest(c, "User ID is required")
		return
	}

	// Get user and profile information
	user, profile, err := GetUserProfileService(c.Request.Context(), userID)
	if err != nil {
		response.NotFound(c, "User profile not found")
		return
	}

	// Prepare response data (public information only)
	userData := map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"full_name":  user.FirstName + " " + user.LastName,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	}

	// Add prefix information if available (public data)
	if user.Prefix != nil {
		userData["prefix"] = map[string]interface{}{
			"id":           user.Prefix.ID,
			"code":         user.Prefix.Code,
			"name_th":      user.Prefix.NameTH,
			"name_en":      user.Prefix.NameEN,
			"abbreviation": user.Prefix.Abbreviation,
		}
	}

	// Add gender information if available (public data)
	if user.Gender != nil {
		userData["gender"] = map[string]interface{}{
			"id":           user.Gender.ID,
			"code":         user.Gender.Code,
			"name_th":      user.Gender.NameTH,
			"name_en":      user.Gender.NameEN,
			"abbreviation": user.Gender.Abbreviation,
		}
	}

	// Add school information if available
	if user.School != nil {
		userData["school"] = map[string]interface{}{
			"id":   user.School.ID,
			"name": user.School.Name,
		}
	}

	// Add public profile information if available
	if profile != nil {
		publicProfile := map[string]interface{}{
			"id":              profile.ID,
			"user_id":         profile.UserID,
			"full_name":       profile.FullName,
			"bio":             profile.Bio,
			"website":         profile.Website,
			"profile_picture": profile.ProfilePicture,
			"city":            profile.City,
			"country":         profile.Country,
		}
		userData["profile"] = publicProfile
	}

	response.Success(c, userData)
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// TODO: Implement logout logic (invalidate token)
	response.Success(c, gin.H{
		"message": "Logout functionality to be implemented",
	})
}

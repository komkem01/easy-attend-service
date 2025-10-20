package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	config "github.com/komkem01/easy-attend-service/configs"
	"github.com/komkem01/easy-attend-service/model"
	"github.com/komkem01/easy-attend-service/requests"
	"github.com/komkem01/easy-attend-service/utils"
	"github.com/komkem01/easy-attend-service/utils/jwt"
	"github.com/uptrace/bun"
)

var db = config.Database()

// FindOrCreateSchoolService finds existing school by name or creates a new one
func FindOrCreateSchoolService(ctx context.Context, schoolName string) (*model.Schools, error) {
	school := &model.Schools{}

	// Try to find existing school by name (case-insensitive)
	err := db.NewSelect().Model(school).
		Where("LOWER(name) = LOWER(?)", schoolName).
		Where("is_active = true").
		Scan(ctx)

	if err == nil {
		// School found, return it
		return school, nil
	}

	// School not found, create new one
	newSchool := &model.Schools{
		ID:        uuid.New(),
		Name:      schoolName,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = db.NewInsert().Model(newSchool).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return newSchool, nil
}

func LoginUserService(ctx context.Context, req requests.LoginRequest) (*model.Users, error) {
	ex, err := db.NewSelect().TableExpr("users").Where("email = ?", req.Email).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("email or password not found")
	}

	user := &model.Users{}

	err = db.NewSelect().Model(user).
		Relation("School").
		Relation("Prefix").
		Relation("Gender").
		Where("u.email = ?", req.Email).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	// Check password using utils function
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.New("email or password not found")
	}

	// Update last login time
	user.LastLoginAt = &time.Time{}
	*user.LastLoginAt = time.Now()
	_, err = db.NewUpdate().Model(user).Column("last_login_at").Where("id = ?", user.ID).Exec(ctx)
	if err != nil {
		// Log error but don't fail login
		// log.Printf("Failed to update last login time: %v", err)
	}

	return user, nil
}

// RegisterUserService creates a new user account
func RegisterUserService(ctx context.Context, req requests.RegisterRequest) (*model.Users, error) {
	// Validate password confirmation
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password confirmation does not match")
	}

	// Check if email already exists
	exists, err := db.NewSelect().TableExpr("users").Where("email = ?", req.Email).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("email already registered")
	}

	// Check if username already exists
	usernameExists, err := db.NewSelect().TableExpr("users").Where("username = ?", req.Username).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if usernameExists {
		return nil, errors.New("username already taken")
	}

	// Find or create school
	school, err := FindOrCreateSchoolService(ctx, req.SchoolName)
	if err != nil {
		return nil, errors.New("failed to process school information")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// Find gender ID if gender name is provided
	var genderID *int
	if req.Gender != nil {
		genderID, err = FindGenderIDByName(ctx, *req.Gender)
		if err != nil {
			// If gender not found, ignore the error and continue without gender
			genderID = nil
		}
	}

	// Find prefix ID if prefix name is provided
	var prefixID *int
	if req.Prefix != nil {
		prefixID, err = FindPrefixIDByName(ctx, *req.Prefix)
		if err != nil {
			// If prefix not found, ignore the error and continue without prefix
			prefixID = nil
		}
	}

	// Create new user
	user := &model.Users{
		ID:            uuid.New(),
		Username:      req.Username,
		Email:         req.Email,
		PasswordHash:  hashedPassword,
		PrefixID:      prefixID,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		GenderID:      genderID,
		Role:          req.UserType,
		SchoolID:      &school.ID,
		Phone:         req.Phone,
		IsActive:      true,
		EmailVerified: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Insert user into database
	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// Load school information for response
	user.School = school

	return user, nil
}

// AuthResponseService generates JWT tokens for authenticated user
func AuthResponseService(user *model.Users) (map[string]interface{}, error) {
	// Generate access token
	accessToken, err := jwt.GenerateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	// Generate refresh token
	refreshToken, err := jwt.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	userInfo := map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"full_name":  user.FirstName + " " + user.LastName,
		"role":       user.Role,
		"school_id":  user.SchoolID,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt,
	}

	// Add prefix information if available
	if user.Prefix != nil {
		userInfo["prefix"] = map[string]interface{}{
			"id":           user.Prefix.ID,
			"code":         user.Prefix.Code,
			"name_th":      user.Prefix.NameTH,
			"name_en":      user.Prefix.NameEN,
			"abbreviation": user.Prefix.Abbreviation,
		}
	}

	// Add gender information if available
	if user.Gender != nil {
		userInfo["gender"] = map[string]interface{}{
			"id":           user.Gender.ID,
			"code":         user.Gender.Code,
			"name_th":      user.Gender.NameTH,
			"name_en":      user.Gender.NameEN,
			"abbreviation": user.Gender.Abbreviation,
		}
	}

	// Add school information if available
	if user.School != nil {
		userInfo["school"] = map[string]interface{}{
			"id":   user.School.ID,
			"name": user.School.Name,
		}
	}

	response := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
		"expires_in":    jwt.GetAccessTokenExpiry(),
		"user":          userInfo,
	}

	return response, nil
}

// RefreshTokenService generates new access token from refresh token
func RefreshTokenService(ctx context.Context, refreshToken string) (map[string]interface{}, error) {
	// Validate refresh token
	claims, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Get user from database
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	user := &model.Users{}
	err = db.NewSelect().Model(user).Where("id = ? AND is_active = true", userID).Scan(ctx)
	if err != nil {
		return nil, errors.New("user not found or inactive")
	}

	// Generate new access token
	accessToken, err := jwt.GenerateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	response := map[string]interface{}{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   jwt.GetAccessTokenExpiry(),
	}

	return response, nil
}

// UpdateUserDataService updates user table data
func UpdateUserDataService(ctx context.Context, userID string, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil // No data to update
	}

	// Parse user ID
	uid, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	// Add updated_at timestamp
	updateData["updated_at"] = time.Now()

	// Build and execute update query
	query := db.NewUpdate().Model((*model.Users)(nil)).Where("id = ?", uid)

	for key, value := range updateData {
		query = query.Set("? = ?", bun.Ident(key), value)
	}

	_, err = query.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

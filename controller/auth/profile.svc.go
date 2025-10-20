package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/model"
)

// GetUserProfileService retrieves user profile information
func GetUserProfileService(ctx context.Context, userID string) (*model.Users, *model.UserProfiles, error) {
	// Parse user ID
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, nil, errors.New("invalid user ID")
	}

	// Get user basic information
	user := &model.Users{}
	err = db.NewSelect().
		Model(user).
		Relation("School").
		Relation("Prefix").
		Relation("Gender").
		Where("u.id = ? AND u.is_active = true", uid).
		Scan(ctx)
	if err != nil {
		return nil, nil, errors.New("user not found")
	}

	// Get user profile (optional - may not exist)
	profile := &model.UserProfiles{}
	err = db.NewSelect().
		Model(profile).
		Where("user_id = ?", uid).
		Scan(ctx)
	if err != nil {
		// Profile doesn't exist, create empty one
		profile = nil
	}

	return user, profile, nil
}

// UpdateUserProfileService updates user profile information
func UpdateUserProfileService(ctx context.Context, userID string, profileData map[string]interface{}) (*model.UserProfiles, error) {
	// Parse user ID
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Check if user exists
	userExists, err := db.NewSelect().
		TableExpr("users").
		Where("id = ?", uid).
		Exists(ctx)
	if err != nil || !userExists {
		return nil, errors.New("user not found")
	}

	// Check if profile exists
	profile := &model.UserProfiles{}
	err = db.NewSelect().
		Model(profile).
		Where("user_id = ?", uid).
		Scan(ctx)

	if err != nil {
		// Profile doesn't exist, create new one
		profile = &model.UserProfiles{
			ID:        uuid.New(),
			UserID:    uid,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Set profile data
		setProfileFields(profile, profileData)

		// Insert new profile
		_, err = db.NewInsert().Model(profile).Exec(ctx)
		if err != nil {
			return nil, errors.New("failed to create user profile")
		}
	} else {
		// Profile exists, update it
		setProfileFields(profile, profileData)
		profile.UpdatedAt = time.Now()

		// Update existing profile
		_, err = db.NewUpdate().
			Model(profile).
			Where("user_id = ?", uid).
			Exec(ctx)
		if err != nil {
			return nil, errors.New("failed to update user profile")
		}
	}

	return profile, nil
}

// setProfileFields sets profile fields from map data
func setProfileFields(profile *model.UserProfiles, data map[string]interface{}) {
	if val, ok := data["first_name"].(string); ok && val != "" {
		profile.FirstName = &val
	}
	if val, ok := data["last_name"].(string); ok && val != "" {
		profile.LastName = &val
	}
	if val, ok := data["full_name"].(string); ok && val != "" {
		profile.FullName = &val
	}
	if val, ok := data["phone_number"].(string); ok && val != "" {
		profile.PhoneNumber = &val
	}
	if val, ok := data["address"].(string); ok && val != "" {
		profile.Address = &val
	}
	if val, ok := data["city"].(string); ok && val != "" {
		profile.City = &val
	}
	if val, ok := data["state"].(string); ok && val != "" {
		profile.State = &val
	}
	if val, ok := data["postal_code"].(string); ok && val != "" {
		profile.PostalCode = &val
	}
	if val, ok := data["country"].(string); ok && val != "" {
		profile.Country = &val
	}
	if val, ok := data["bio"].(string); ok && val != "" {
		profile.Bio = &val
	}
	if val, ok := data["website"].(string); ok && val != "" {
		profile.Website = &val
	}
	if val, ok := data["gender"].(string); ok && val != "" {
		profile.Gender = &val
	}
	if val, ok := data["profile_picture"].(string); ok && val != "" {
		profile.ProfilePicture = &val
	}

	// Handle date of birth
	if val, ok := data["date_of_birth"].(string); ok && val != "" {
		if dob, err := time.Parse("2006-01-02", val); err == nil {
			profile.DateOfBirth = &dob
		}
	}
}

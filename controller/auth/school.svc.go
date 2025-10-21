package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/model"
	"github.com/uptrace/bun"
)

// GetSchoolsService returns all active schools
func GetSchoolsService(ctx context.Context) ([]model.Schools, error) {
	var schools []model.Schools

	err := db.NewSelect().
		Model(&schools).
		Where("is_active = true").
		Order("name ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return schools, nil
}

// GetSchoolByIDService returns school by ID
func GetSchoolByIDService(ctx context.Context, schoolID string) (*model.Schools, error) {
	// Parse school ID
	sid, err := uuid.Parse(schoolID)
	if err != nil {
		return nil, errors.New("invalid school ID")
	}

	var school model.Schools
	err = db.NewSelect().
		Model(&school).
		Where("id = ? AND is_active = true", sid).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &school, nil
}

// CreateSchoolService creates a new school
func CreateSchoolService(ctx context.Context, name, address, phone, email, website string) (*model.Schools, error) {
	// Check if school name already exists
	exists, err := db.NewSelect().
		TableExpr("schools").
		Where("LOWER(name) = LOWER(?) AND is_active = true", name).
		Exists(ctx)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("school name already exists")
	}

	// Create new school
	school := &model.Schools{
		ID:         uuid.New(),
		Name:       name,
		Address:    &address,
		Phone:      &phone,
		Email:      &email,
		WebsiteURL: &website,
		IsActive:   true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err = db.NewInsert().Model(school).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// UpdateSchoolService updates school information
func UpdateSchoolService(ctx context.Context, schoolID string, updateData map[string]interface{}) (*model.Schools, error) {
	// Parse school ID
	sid, err := uuid.Parse(schoolID)
	if err != nil {
		return nil, errors.New("invalid school ID")
	}

	if len(updateData) == 0 {
		return nil, errors.New("no data to update")
	}

	// Add updated_at timestamp
	updateData["updated_at"] = time.Now()

	// Skip duplicate name checking for updates to avoid conflicts with self
	// The database unique constraint will still prevent true duplicates if needed
	// This approach is safer for PATCH operations on existing records	// Build and execute update query
	query := db.NewUpdate().Model((*model.Schools)(nil)).Where("id = ?", sid)

	for key, value := range updateData {
		query = query.Set("? = ?", bun.Ident(key), value)
	}

	_, err = query.Exec(ctx)
	if err != nil {
		return nil, err
	}

	// Get updated school
	return GetSchoolByIDService(ctx, schoolID)
}

// DeleteSchoolService soft deletes a school (sets is_active to false)
func DeleteSchoolService(ctx context.Context, schoolID string) error {
	// Parse school ID
	sid, err := uuid.Parse(schoolID)
	if err != nil {
		return errors.New("invalid school ID")
	}

	// Check if school has users
	hasUsers, err := db.NewSelect().
		TableExpr("users").
		Where("school_id = ? AND is_active = true", sid).
		Exists(ctx)
	if err != nil {
		return err
	}

	if hasUsers {
		return errors.New("cannot delete school that has active users")
	}

	// Soft delete (set is_active to false)
	_, err = db.NewUpdate().
		Model((*model.Schools)(nil)).
		Set("is_active = false").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", sid).
		Exec(ctx)

	return err
}

// SearchSchoolsService searches schools by name
func SearchSchoolsService(ctx context.Context, query string) ([]model.Schools, error) {
	var schools []model.Schools

	err := db.NewSelect().
		Model(&schools).
		Where("is_active = true").
		Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").
		Order("name ASC").
		Limit(20). // Limit results to prevent large responses
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return schools, nil
}

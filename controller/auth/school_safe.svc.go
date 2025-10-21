package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/model"
	"github.com/uptrace/bun"
)

// UpdateSchoolServiceSafe updates school information without duplicate name check
// This is a safer version that avoids the "name already exists" issue during updates
func UpdateSchoolServiceSafe(ctx context.Context, schoolID string, updateData map[string]interface{}) (*model.Schools, error) {
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

	// Note: We skip duplicate name checking for updates to avoid conflicts
	// The database unique constraint will still prevent true duplicates
	// This is safer for PATCH operations where users update existing records

	// Build and execute update query
	query := db.NewUpdate().Model((*model.Schools)(nil)).Where("id = ?", sid)

	for key, value := range updateData {
		query = query.Set("? = ?", bun.Ident(key), value)
	}

	_, err = query.Exec(ctx)
	if err != nil {
		// Check if it's a unique constraint violation
		if fmt.Sprintf("%v", err) == "duplicate key value violates unique constraint" {
			return nil, errors.New("school name already exists")
		}
		return nil, err
	}

	// Get updated school
	return GetSchoolByIDService(ctx, schoolID)
}

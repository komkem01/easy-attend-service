package auth

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/model"
	"github.com/komkem01/easy-attend-service/requests"
	"github.com/uptrace/bun"
)

// ClassroomService handles classroom business logic
type ClassroomService struct {
	db *bun.DB
}

// NewClassroomService creates a new classroom service
func NewClassroomService(db *bun.DB) *ClassroomService {
	return &ClassroomService{db: db}
}

// CreateClassroomService creates a new classroom
func (s *ClassroomService) CreateClassroomService(ctx context.Context, req *requests.CreateClassroomRequest, teacherID uuid.UUID) (*model.Classrooms, error) {
	// Generate unique classroom code
	classroomCode, err := s.generateClassroomCode(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate classroom code: %w", err)
	}

	// Set default max students if not provided
	maxStudents := 50
	if req.MaxStudents != nil {
		maxStudents = *req.MaxStudents
	}

	classroom := &model.Classrooms{
		ID:            uuid.New(),
		SchoolID:      req.SchoolID,
		Name:          req.Name,
		Subject:       req.Subject,
		Description:   req.Description,
		GradeLevel:    req.GradeLevel,
		Section:       req.Section,
		RoomNumber:    req.RoomNumber,
		TeacherID:     teacherID,
		ClassroomCode: classroomCode,
		MaxStudents:   maxStudents,
		Schedule:      req.Schedule,
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = s.db.NewInsert().Model(classroom).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create classroom: %w", err)
	}

	// Load relationships
	err = s.db.NewSelect().
		Model(classroom).
		Relation("School").
		Relation("Teacher").
		Where("c.id = ?", classroom.ID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load classroom relationships: %w", err)
	}

	return classroom, nil
}

// GetClassroomsService retrieves classrooms with filtering and pagination
func (s *ClassroomService) GetClassroomsService(ctx context.Context, req *requests.ClassroomQueryRequest) ([]*model.Classrooms, int64, error) {
	query := s.db.NewSelect().
		Model((*model.Classrooms)(nil)).
		Relation("School").
		Relation("Teacher").
		Where("c.deleted_at IS NULL")

	// Apply filters
	if req.Search != nil && *req.Search != "" {
		searchTerm := "%" + strings.ToLower(*req.Search) + "%"
		query = query.Where("(LOWER(c.name) LIKE ? OR LOWER(c.subject) LIKE ? OR LOWER(c.classroom_code) LIKE ?)",
			searchTerm, searchTerm, searchTerm)
	}

	if req.SchoolID != nil {
		query = query.Where("c.school_id = ?", *req.SchoolID)
	}

	if req.TeacherID != nil {
		query = query.Where("c.teacher_id = ?", *req.TeacherID)
	}

	if req.Subject != nil && *req.Subject != "" {
		query = query.Where("LOWER(c.subject) = LOWER(?)", *req.Subject)
	}

	if req.GradeLevel != nil && *req.GradeLevel != "" {
		query = query.Where("c.grade_level = ?", *req.GradeLevel)
	}

	if req.IsActive != nil {
		query = query.Where("c.is_active = ?", *req.IsActive)
	}

	// Count total records
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count classrooms: %w", err)
	}

	// Apply pagination
	page := 1
	limit := 20
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 {
		limit = req.Limit
	}

	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	// Order by created_at desc
	query = query.Order("c.created_at DESC")

	var classrooms []*model.Classrooms
	err = query.Scan(ctx, &classrooms)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve classrooms: %w", err)
	}

	return classrooms, int64(total), nil
}

// GetClassroomByIDService retrieves a classroom by ID
func (s *ClassroomService) GetClassroomByIDService(ctx context.Context, id uuid.UUID) (*model.Classrooms, error) {
	var classroom model.Classrooms
	err := s.db.NewSelect().
		Model(&classroom).
		Relation("School").
		Relation("Teacher").
		Relation("ClassroomStudents").
		Relation("ClassroomStudents.Student").
		Where("c.id = ? AND c.deleted_at IS NULL", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("classroom not found")
		}
		return nil, fmt.Errorf("failed to retrieve classroom: %w", err)
	}

	return &classroom, nil
}

// GetClassroomByCodeService retrieves a classroom by classroom code
func (s *ClassroomService) GetClassroomByCodeService(ctx context.Context, code string) (*model.Classrooms, error) {
	var classroom model.Classrooms
	err := s.db.NewSelect().
		Model(&classroom).
		Relation("School").
		Relation("Teacher").
		Where("c.classroom_code = ? AND c.deleted_at IS NULL", code).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("classroom not found")
		}
		return nil, fmt.Errorf("failed to retrieve classroom: %w", err)
	}

	return &classroom, nil
}

// UpdateClassroomService updates a classroom
func (s *ClassroomService) UpdateClassroomService(ctx context.Context, id uuid.UUID, req *requests.UpdateClassroomRequest, teacherID uuid.UUID) (*model.Classrooms, error) {
	// Check if classroom exists and user has permission
	var classroom model.Classrooms
	err := s.db.NewSelect().
		Model(&classroom).
		Where("c.id = ? AND c.deleted_at IS NULL", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("classroom not found")
		}
		return nil, fmt.Errorf("failed to retrieve classroom: %w", err)
	}

	// Check if user is the teacher of this classroom (or admin)
	if classroom.TeacherID != teacherID {
		// TODO: Add admin role check here
		return nil, fmt.Errorf("unauthorized to update this classroom")
	}

	// Build update data
	updateData := make(map[string]interface{})

	if req.SchoolID != nil {
		updateData["school_id"] = *req.SchoolID
	}
	if req.Name != nil {
		updateData["name"] = *req.Name
	}
	if req.Subject != nil {
		updateData["subject"] = *req.Subject
	}
	if req.Description != nil {
		updateData["description"] = *req.Description
	}
	if req.GradeLevel != nil {
		updateData["grade_level"] = *req.GradeLevel
	}
	if req.Section != nil {
		updateData["section"] = *req.Section
	}
	if req.RoomNumber != nil {
		updateData["room_number"] = *req.RoomNumber
	}
	if req.MaxStudents != nil {
		updateData["max_students"] = *req.MaxStudents
	}
	if req.Schedule != nil {
		updateData["schedule"] = *req.Schedule
	}
	if req.IsActive != nil {
		updateData["is_active"] = *req.IsActive
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no data to update")
	}

	updateData["updated_at"] = time.Now()

	// Apply updates dynamically
	query := s.db.NewUpdate().Model(&classroom).Where("id = ?", id)
	for key, value := range updateData {
		query = query.Set("? = ?", bun.Ident(key), value)
	}

	_, err = query.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update classroom: %w", err)
	}

	// Retrieve updated classroom with relationships
	err = s.db.NewSelect().
		Model(&classroom).
		Relation("School").
		Relation("Teacher").
		Where("c.id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load updated classroom: %w", err)
	}

	return &classroom, nil
}

// DeleteClassroomService soft deletes a classroom
func (s *ClassroomService) DeleteClassroomService(ctx context.Context, id uuid.UUID, teacherID uuid.UUID) error {
	// Check if classroom exists and user has permission
	var classroom model.Classrooms
	err := s.db.NewSelect().
		Model(&classroom).
		Where("c.id = ? AND c.deleted_at IS NULL", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("classroom not found")
		}
		return fmt.Errorf("failed to retrieve classroom: %w", err)
	}

	// Check if user is the teacher of this classroom (or admin)
	if classroom.TeacherID != teacherID {
		// TODO: Add admin role check here
		return fmt.Errorf("unauthorized to delete this classroom")
	}

	// Check if classroom has active students
	var studentCount int
	studentCount, err = s.db.NewSelect().
		Model((*model.ClassroomStudents)(nil)).
		Where("classroom_id = ? AND is_active = true", id).
		Count(ctx)

	if err != nil {
		return fmt.Errorf("failed to check classroom students: %w", err)
	}

	if studentCount > 0 {
		return fmt.Errorf("cannot delete classroom with active students")
	}

	// Soft delete the classroom
	_, err = s.db.NewDelete().Model(&classroom).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete classroom: %w", err)
	}

	return nil
}

// generateClassroomCode generates a unique classroom code
func (s *ClassroomService) generateClassroomCode(ctx context.Context) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	for attempts := 0; attempts < 10; attempts++ {
		// Generate random code
		code := make([]byte, codeLength)
		for i := range code {
			code[i] = charset[rand.Intn(len(charset))]
		}

		codeStr := string(code)

		// Check if code already exists
		var count int
		count, err := s.db.NewSelect().
			Model((*model.Classrooms)(nil)).
			Where("classroom_code = ?", codeStr).
			Count(ctx)

		if err != nil {
			return "", fmt.Errorf("failed to check classroom code uniqueness: %w", err)
		}

		if count == 0 {
			return codeStr, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique classroom code after 10 attempts")
}

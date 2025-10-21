package auth

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/model"
	"github.com/komkem01/easy-attend-service/requests"
	"github.com/uptrace/bun"
)

// AssignmentService handles assignment business logic
type AssignmentService struct {
	db *bun.DB
}

// NewAssignmentService creates a new assignment service
func NewAssignmentService(db *bun.DB) *AssignmentService {
	return &AssignmentService{db: db}
}

// CreateAssignmentService creates a new assignment
func (s *AssignmentService) CreateAssignmentService(ctx context.Context, req *requests.CreateAssignmentRequest, teacherID uuid.UUID) (*model.Assignments, error) {
	// Verify that the teacher has access to this classroom
	err := s.verifyClassroomAccess(ctx, req.ClassroomID, teacherID)
	if err != nil {
		return nil, err
	}

	// Set default values
	assignmentType := "homework"
	if req.AssignmentType != nil {
		assignmentType = *req.AssignmentType
	}

	maxScore := 100.0
	if req.MaxScore != nil {
		maxScore = *req.MaxScore
	}

	weight := 1.0
	if req.Weight != nil {
		weight = *req.Weight
	}

	allowLateSubmission := false
	if req.AllowLateSubmission != nil {
		allowLateSubmission = *req.AllowLateSubmission
	}

	latePenaltyPercent := 0.0
	if req.LatePenaltyPercent != nil {
		latePenaltyPercent = *req.LatePenaltyPercent
	}

	submissionFormat := "both"
	if req.SubmissionFormat != nil {
		submissionFormat = *req.SubmissionFormat
	}

	maxFileSizeMB := 10
	if req.MaxFileSizeMB != nil {
		maxFileSizeMB = *req.MaxFileSizeMB
	}

	isPublished := false
	if req.IsPublished != nil {
		isPublished = *req.IsPublished
	}

	assignment := &model.Assignments{
		ID:                  uuid.New(),
		ClassroomID:         req.ClassroomID,
		Title:               req.Title,
		Description:         req.Description,
		Instructions:        req.Instructions,
		AssignmentType:      assignmentType,
		DueDate:             req.DueDate,
		MaxScore:            maxScore,
		Weight:              weight,
		AllowLateSubmission: allowLateSubmission,
		LatePenaltyPercent:  latePenaltyPercent,
		SubmissionFormat:    submissionFormat,
		MaxFileSizeMB:       maxFileSizeMB,
		AllowedFileTypes:    req.AllowedFileTypes,
		IsPublished:         isPublished,
		Status:              "draft",
		CreatedBy:           teacherID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	_, err = s.db.NewInsert().Model(assignment).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create assignment: %w", err)
	}

	// Load relationships
	err = s.db.NewSelect().
		Model(assignment).
		Relation("Classroom").
		Relation("Creator").
		Where("a.id = ?", assignment.ID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load assignment relationships: %w", err)
	}

	return assignment, nil
}

// GetAssignmentsService retrieves assignments with filtering and pagination
func (s *AssignmentService) GetAssignmentsService(ctx context.Context, req *requests.AssignmentQueryRequest) ([]*model.Assignments, int64, error) {
	query := s.db.NewSelect().
		Model((*model.Assignments)(nil)).
		Relation("Classroom").
		Relation("Creator").
		Where("a.deleted_at IS NULL")

	// Apply filters
	if req.Search != nil && *req.Search != "" {
		searchTerm := "%" + strings.ToLower(*req.Search) + "%"
		query = query.Where("(LOWER(a.title) LIKE ? OR LOWER(a.description) LIKE ?)",
			searchTerm, searchTerm)
	}

	if req.ClassroomID != nil {
		query = query.Where("a.classroom_id = ?", *req.ClassroomID)
	}

	if req.CreatedBy != nil {
		query = query.Where("a.created_by = ?", *req.CreatedBy)
	}

	if req.AssignmentType != nil && *req.AssignmentType != "" {
		query = query.Where("a.assignment_type = ?", *req.AssignmentType)
	}

	if req.Status != nil && *req.Status != "" {
		query = query.Where("a.status = ?", *req.Status)
	}

	if req.IsPublished != nil {
		query = query.Where("a.is_published = ?", *req.IsPublished)
	}

	if req.DueSoon != nil && *req.DueSoon {
		// Due within 7 days
		weekFromNow := time.Now().AddDate(0, 0, 7)
		query = query.Where("a.due_date IS NOT NULL AND a.due_date <= ? AND a.due_date >= ?", weekFromNow, time.Now())
	}

	if req.Overdue != nil && *req.Overdue {
		// Past due date
		query = query.Where("a.due_date IS NOT NULL AND a.due_date < ?", time.Now())
	}

	// Count total records
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count assignments: %w", err)
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
	query = query.Order("a.created_at DESC")

	var assignments []*model.Assignments
	err = query.Scan(ctx, &assignments)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve assignments: %w", err)
	}

	return assignments, int64(total), nil
}

// GetAssignmentByIDService retrieves an assignment by ID
func (s *AssignmentService) GetAssignmentByIDService(ctx context.Context, id uuid.UUID) (*model.Assignments, error) {
	var assignment model.Assignments
	err := s.db.NewSelect().
		Model(&assignment).
		Relation("Classroom").
		Relation("Creator").
		Where("a.id = ? AND a.deleted_at IS NULL", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("assignment not found")
		}
		return nil, fmt.Errorf("failed to retrieve assignment: %w", err)
	}

	return &assignment, nil
}

// UpdateAssignmentService updates an assignment
func (s *AssignmentService) UpdateAssignmentService(ctx context.Context, id uuid.UUID, req *requests.UpdateAssignmentRequest, teacherID uuid.UUID) (*model.Assignments, error) {
	// Check if assignment exists and user has permission
	var assignment model.Assignments
	err := s.db.NewSelect().
		Model(&assignment).
		Relation("Classroom").
		Where("a.id = ? AND a.deleted_at IS NULL", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("assignment not found")
		}
		return nil, fmt.Errorf("failed to retrieve assignment: %w", err)
	}

	// Check if user is the creator of this assignment or teacher of the classroom
	if assignment.CreatedBy != teacherID && assignment.Classroom.TeacherID != teacherID {
		return nil, fmt.Errorf("unauthorized to update this assignment")
	}

	// Build update data
	updateData := make(map[string]interface{})

	if req.Title != nil {
		updateData["title"] = *req.Title
	}
	if req.Description != nil {
		updateData["description"] = *req.Description
	}
	if req.Instructions != nil {
		updateData["instructions"] = *req.Instructions
	}
	if req.AssignmentType != nil {
		updateData["assignment_type"] = *req.AssignmentType
	}
	if req.DueDate != nil {
		updateData["due_date"] = *req.DueDate
	}
	if req.MaxScore != nil {
		updateData["max_score"] = *req.MaxScore
	}
	if req.Weight != nil {
		updateData["weight"] = *req.Weight
	}
	if req.AllowLateSubmission != nil {
		updateData["allow_late_submission"] = *req.AllowLateSubmission
	}
	if req.LatePenaltyPercent != nil {
		updateData["late_penalty_percent"] = *req.LatePenaltyPercent
	}
	if req.SubmissionFormat != nil {
		updateData["submission_format"] = *req.SubmissionFormat
	}
	if req.MaxFileSizeMB != nil {
		updateData["max_file_size_mb"] = *req.MaxFileSizeMB
	}
	if req.AllowedFileTypes != nil {
		updateData["allowed_file_types"] = *req.AllowedFileTypes
	}
	if req.IsPublished != nil {
		updateData["is_published"] = *req.IsPublished
	}
	if req.Status != nil {
		updateData["status"] = *req.Status
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no data to update")
	}

	updateData["updated_at"] = time.Now()

	// Apply updates dynamically
	query := s.db.NewUpdate().Model(&assignment).Where("id = ?", id)
	for key, value := range updateData {
		query = query.Set("? = ?", bun.Ident(key), value)
	}

	_, err = query.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update assignment: %w", err)
	}

	// Retrieve updated assignment with relationships
	err = s.db.NewSelect().
		Model(&assignment).
		Relation("Classroom").
		Relation("Creator").
		Where("a.id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load updated assignment: %w", err)
	}

	return &assignment, nil
}

// DeleteAssignmentService soft deletes an assignment
func (s *AssignmentService) DeleteAssignmentService(ctx context.Context, id uuid.UUID, teacherID uuid.UUID) error {
	// Check if assignment exists and user has permission
	var assignment model.Assignments
	err := s.db.NewSelect().
		Model(&assignment).
		Relation("Classroom").
		Where("a.id = ? AND a.deleted_at IS NULL", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("assignment not found")
		}
		return fmt.Errorf("failed to retrieve assignment: %w", err)
	}

	// Check if user is the creator of this assignment or teacher of the classroom
	if assignment.CreatedBy != teacherID && assignment.Classroom.TeacherID != teacherID {
		return fmt.Errorf("unauthorized to delete this assignment")
	}

	// Check if assignment has submissions
	var submissionCount int
	submissionCount, err = s.db.NewSelect().
		Model((*model.AssignmentSubmissions)(nil)).
		Where("assignment_id = ? AND deleted_at IS NULL", id).
		Count(ctx)

	if err != nil {
		return fmt.Errorf("failed to check assignment submissions: %w", err)
	}

	if submissionCount > 0 {
		return fmt.Errorf("cannot delete assignment with existing submissions")
	}

	// Soft delete the assignment
	_, err = s.db.NewDelete().Model(&assignment).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete assignment: %w", err)
	}

	return nil
}

// PublishAssignmentService publishes an assignment
func (s *AssignmentService) PublishAssignmentService(ctx context.Context, id uuid.UUID, teacherID uuid.UUID) (*model.Assignments, error) {
	// Update assignment to published status
	req := &requests.UpdateAssignmentRequest{
		IsPublished: &[]bool{true}[0],
		Status:      &[]string{"published"}[0],
	}

	return s.UpdateAssignmentService(ctx, id, req, teacherID)
}

// verifyClassroomAccess checks if the teacher has access to the classroom
func (s *AssignmentService) verifyClassroomAccess(ctx context.Context, classroomID, teacherID uuid.UUID) error {
	var classroom model.Classrooms
	err := s.db.NewSelect().
		Model(&classroom).
		Where("c.id = ? AND c.teacher_id = ? AND c.deleted_at IS NULL", classroomID, teacherID).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("classroom not found or access denied")
		}
		return fmt.Errorf("failed to verify classroom access: %w", err)
	}

	return nil
}

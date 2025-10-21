package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/requests"
	"github.com/komkem01/easy-attend-service/response"
	"github.com/uptrace/bun"
)

// Helper function to get UUID from JWT context
func GetAssignmentUserUUIDFromContext(c *gin.Context) (uuid.UUID, error) {
	teacherID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, fmt.Errorf("user not authenticated")
	}

	teacherIDStr, ok := teacherID.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user ID format")
	}

	teacherUUID, err := uuid.Parse(teacherIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID format")
	}

	return teacherUUID, nil
}

// AssignmentController handles assignment HTTP requests
type AssignmentController struct {
	assignmentService *AssignmentService
}

// NewAssignmentController creates a new assignment controller
func NewAssignmentController(db *bun.DB) *AssignmentController {
	return &AssignmentController{
		assignmentService: NewAssignmentService(db),
	}
}

// CreateAssignment creates a new assignment
func (ctrl *AssignmentController) CreateAssignment(c *gin.Context) {
	var req requests.CreateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetAssignmentUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	assignment, err := ctrl.assignmentService.CreateAssignmentService(c.Request.Context(), &req, teacherUUID)
	if err != nil {
		if err.Error() == "classroom not found or access denied" {
			response.Forbidden(c, "You don't have access to this classroom")
			return
		}
		response.InternalServerError(c, "Failed to create assignment: "+err.Error())
		return
	}

	response.Created(c, assignment)
}

// GetAssignments retrieves assignments with filtering and pagination
func (ctrl *AssignmentController) GetAssignments(c *gin.Context) {
	var req requests.AssignmentQueryRequest

	// Parse query parameters
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	if search := c.Query("search"); search != "" {
		req.Search = &search
	}

	if classroomIDStr := c.Query("classroom_id"); classroomIDStr != "" {
		if classroomID, err := uuid.Parse(classroomIDStr); err == nil {
			req.ClassroomID = &classroomID
		}
	}

	if createdByStr := c.Query("created_by"); createdByStr != "" {
		if createdBy, err := uuid.Parse(createdByStr); err == nil {
			req.CreatedBy = &createdBy
		}
	}

	if assignmentType := c.Query("assignment_type"); assignmentType != "" {
		req.AssignmentType = &assignmentType
	}

	if status := c.Query("status"); status != "" {
		req.Status = &status
	}

	if isPublishedStr := c.Query("is_published"); isPublishedStr != "" {
		if isPublished, err := strconv.ParseBool(isPublishedStr); err == nil {
			req.IsPublished = &isPublished
		}
	}

	if dueSoonStr := c.Query("due_soon"); dueSoonStr != "" {
		if dueSoon, err := strconv.ParseBool(dueSoonStr); err == nil {
			req.DueSoon = &dueSoon
		}
	}

	if overdueStr := c.Query("overdue"); overdueStr != "" {
		if overdue, err := strconv.ParseBool(overdueStr); err == nil {
			req.Overdue = &overdue
		}
	}

	assignments, total, err := ctrl.assignmentService.GetAssignmentsService(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve assignments: "+err.Error())
		return
	}

	// Calculate pagination info
	page := 1
	limit := 20
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 {
		limit = req.Limit
	}

	totalPages := (int(total) + limit - 1) / limit
	hasNext := page < totalPages
	hasPrev := page > 1

	responseData := gin.H{
		"assignments": assignments,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  totalPages,
			"has_next":     hasNext,
			"has_prev":     hasPrev,
		},
	}

	response.Success(c, responseData)
}

// GetAssignment retrieves a single assignment by ID
func (ctrl *AssignmentController) GetAssignment(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.BadRequest(c, "Assignment ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid assignment ID format")
		return
	}

	assignment, err := ctrl.assignmentService.GetAssignmentByIDService(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "assignment not found" {
			response.NotFound(c, "Assignment not found")
			return
		}
		response.InternalServerError(c, "Failed to retrieve assignment: "+err.Error())
		return
	}

	response.Success(c, assignment)
}

// UpdateAssignment updates an assignment
func (ctrl *AssignmentController) UpdateAssignment(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.BadRequest(c, "Assignment ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid assignment ID format")
		return
	}

	var req requests.UpdateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetAssignmentUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	assignment, err := ctrl.assignmentService.UpdateAssignmentService(c.Request.Context(), id, &req, teacherUUID)
	if err != nil {
		if err.Error() == "assignment not found" {
			response.NotFound(c, "Assignment not found")
			return
		}
		if err.Error() == "unauthorized to update this assignment" {
			response.Forbidden(c, "You don't have permission to update this assignment")
			return
		}
		if err.Error() == "no data to update" {
			response.BadRequest(c, "No data provided for update")
			return
		}
		response.InternalServerError(c, "Failed to update assignment: "+err.Error())
		return
	}

	response.Success(c, assignment)
}

// DeleteAssignment soft deletes an assignment
func (ctrl *AssignmentController) DeleteAssignment(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.BadRequest(c, "Assignment ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid assignment ID format")
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetAssignmentUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	err = ctrl.assignmentService.DeleteAssignmentService(c.Request.Context(), id, teacherUUID)
	if err != nil {
		if err.Error() == "assignment not found" {
			response.NotFound(c, "Assignment not found")
			return
		}
		if err.Error() == "unauthorized to delete this assignment" {
			response.Forbidden(c, "You don't have permission to delete this assignment")
			return
		}
		if err.Error() == "cannot delete assignment with existing submissions" {
			response.BadRequest(c, "Cannot delete assignment with existing submissions")
			return
		}
		response.InternalServerError(c, "Failed to delete assignment: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Assignment deleted successfully",
	})
}

// PublishAssignment publishes an assignment
func (ctrl *AssignmentController) PublishAssignment(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.BadRequest(c, "Assignment ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid assignment ID format")
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetAssignmentUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	assignment, err := ctrl.assignmentService.PublishAssignmentService(c.Request.Context(), id, teacherUUID)
	if err != nil {
		if err.Error() == "assignment not found" {
			response.NotFound(c, "Assignment not found")
			return
		}
		if err.Error() == "unauthorized to update this assignment" {
			response.Forbidden(c, "You don't have permission to publish this assignment")
			return
		}
		response.InternalServerError(c, "Failed to publish assignment: "+err.Error())
		return
	}

	response.Success(c, assignment)
}

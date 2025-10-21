package auth

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/komkem01/easy-attend-service/requests"
	"github.com/komkem01/easy-attend-service/response"
)

type ClassroomController struct {
	classroomService *ClassroomService
}

func NewClassroomController(service *ClassroomService) *ClassroomController {
	return &ClassroomController{
		classroomService: service,
	}
}

// Helper function to get UUID from JWT context
func GetUserUUIDFromContext(c *gin.Context) (uuid.UUID, error) {
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

// CreateClassroom creates a new classroom
func (ctrl *ClassroomController) CreateClassroom(c *gin.Context) {
	var req requests.CreateClassroomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	classroom, err := ctrl.classroomService.CreateClassroomService(c.Request.Context(), &req, teacherUUID)
	if err != nil {
		response.InternalServerError(c, "Failed to create classroom: "+err.Error())
		return
	}

	response.Created(c, map[string]interface{}{
		"id":         classroom.ID,
		"name":       classroom.Name,
		"subject":    classroom.Subject,
		"created_at": classroom.CreatedAt,
	})
}

// GetClassrooms gets all classrooms with optional filtering
func (ctrl *ClassroomController) GetClassrooms(c *gin.Context) {
	var req requests.ClassroomQueryRequest

	// Parse query parameters
	if page := c.Query("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			req.Page = p
		}
	}
	if req.Page == 0 {
		req.Page = 1
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 100 {
			req.Limit = l
		}
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	if search := c.Query("search"); search != "" {
		req.Search = &search
	}
	if subject := c.Query("subject"); subject != "" {
		req.Subject = &subject
	}
	if gradeLevel := c.Query("grade_level"); gradeLevel != "" {
		req.GradeLevel = &gradeLevel
	}

	if teacherIDStr := c.Query("teacher_id"); teacherIDStr != "" {
		if teacherID, err := uuid.Parse(teacherIDStr); err == nil {
			req.TeacherID = &teacherID
		}
	}

	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		if isActive, err := strconv.ParseBool(isActiveStr); err == nil {
			req.IsActive = &isActive
		}
	}

	classrooms, total, err := ctrl.classroomService.GetClassroomsService(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch classrooms: "+err.Error())
		return
	}

	// Calculate pagination
	totalPages := int((total + int64(req.Limit) - 1) / int64(req.Limit))
	hasNext := req.Page < totalPages
	hasPrev := req.Page > 1

	response.Success(c, map[string]interface{}{
		"classrooms": classrooms,
		"pagination": map[string]interface{}{
			"current_page": req.Page,
			"per_page":     req.Limit,
			"total":        total,
			"total_pages":  totalPages,
			"has_next":     hasNext,
			"has_prev":     hasPrev,
		},
	})
}

// GetClassroom gets a single classroom by ID
func (ctrl *ClassroomController) GetClassroom(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid classroom ID format")
		return
	}

	classroom, err := ctrl.classroomService.GetClassroomByIDService(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "classroom not found" {
			response.NotFound(c, "Classroom not found")
		} else {
			response.InternalServerError(c, "Failed to fetch classroom: "+err.Error())
		}
		return
	}

	response.Success(c, classroom)
}

// UpdateClassroom updates an existing classroom
func (ctrl *ClassroomController) UpdateClassroom(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid classroom ID format")
		return
	}

	var req requests.UpdateClassroomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	classroom, err := ctrl.classroomService.UpdateClassroomService(c.Request.Context(), id, &req, teacherUUID)
	if err != nil {
		if err.Error() == "classroom not found" {
			response.NotFound(c, "Classroom not found")
		} else if err.Error() == "access denied: you can only update your own classrooms" {
			response.Forbidden(c, "You can only update your own classrooms")
		} else {
			response.InternalServerError(c, "Failed to update classroom: "+err.Error())
		}
		return
	}

	response.Success(c, map[string]interface{}{
		"id":         classroom.ID,
		"name":       classroom.Name,
		"subject":    classroom.Subject,
		"updated_at": classroom.UpdatedAt,
	})
}

// DeleteClassroom deletes a classroom (soft delete)
func (ctrl *ClassroomController) DeleteClassroom(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid classroom ID format")
		return
	}

	// Get teacher ID from JWT token
	teacherUUID, err := GetUserUUIDFromContext(c)
	if err != nil {
		if err.Error() == "user not authenticated" {
			response.Unauthorized(c, err.Error())
		} else {
			response.InternalServerError(c, err.Error())
		}
		return
	}

	err = ctrl.classroomService.DeleteClassroomService(c.Request.Context(), id, teacherUUID)
	if err != nil {
		if err.Error() == "classroom not found" {
			response.NotFound(c, "Classroom not found")
		} else if err.Error() == "access denied: you can only delete your own classrooms" {
			response.Forbidden(c, "You can only delete your own classrooms")
		} else {
			response.InternalServerError(c, "Failed to delete classroom: "+err.Error())
		}
		return
	}

	c.JSON(200, gin.H{"success": true, "message": "Classroom deleted successfully"})
}

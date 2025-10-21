package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/komkem01/easy-attend-service/model"
)

type StatusResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Status StatusResponse `json:"status"`
	Data   any            `json:"data,omitempty"`
}

type ResponsePaginate struct {
	Status     StatusResponse `json:"status"`
	Data       any            `json:"data,omitempty"`
	Pagination model.Paginate `json:"paginate"`
}

type ResponsePaginate0 struct {
	Status     StatusResponse `json:"status"`
	Data       any            `json:"data,omitempty"`
	Pagination any            `json:"pagination"`
}

func SuccessWithPaginate(ctx *gin.Context, data any, pagination model.Paginate) {
	if pagination.Total == 0 {
		ctx.JSON(http.StatusOK, ResponsePaginate0{
			Status: StatusResponse{
				Code:    200,
				Message: "Success",
			},
			Data:       []any{},
			Pagination: gin.H{},
		})
		return
	} else {
		ctx.JSON(http.StatusOK, ResponsePaginate{
			Status: StatusResponse{
				Code:    200,
				Message: "Success",
			},
			Data:       data,
			Pagination: pagination,
		})
	}
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{StatusResponse{
		Code:    200,
		Message: "Success",
	}, data})
}

// Created sends a 201 Created response
func Created(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusCreated, Response{StatusResponse{
		Code:    201,
		Message: "Created",
	}, data})
}

// InternalError ส่งผลลัพธ์เมื่อมีข้อผิดพลาดภายใน
func InternalError(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusInternalServerError, StatusResponse{
		Code:    500,
		Message: message.(string), // Set the message directly here
	})
}

func NotFound(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusNotFound, StatusResponse{
		Code:    404,
		Message: message.(string), // Set the message directly here
	})
}

// BadRequest ส่งผลลัพธ์เมื่อมีข้อผิดพลาดจากการขอข้อมูลที่ไม่ถูกต้อง
func BadRequest(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusBadRequest, StatusResponse{
		Code:    400,
		Message: message.(string), // Set the message directly here
	})
}

func Unauthorized(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusUnauthorized, StatusResponse{
		Code:    401,
		Message: message.(string),
	})
}

// Forbidden sends a 403 Forbidden response
func Forbidden(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusForbidden, StatusResponse{
		Code:    403,
		Message: message.(string),
	})
}

// Conflict sends a 409 Conflict response
func Conflict(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusConflict, StatusResponse{
		Code:    409,
		Message: message.(string),
	})
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusInternalServerError, StatusResponse{
		Code:    500,
		Message: message.(string),
	})
}

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	// TODO: Get user from JWT token
	response.Success(c, gin.H{
		"message": "Profile functionality to be implemented",
	})
}

// UpdateProfile handles updating user profile (PATCH method for partial updates)
func UpdateProfile(c *gin.Context) {
	var req requests.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// TODO: Implement profile partial update logic using PATCH method
	response.Success(c, gin.H{
		"message": "Profile partial update functionality to be implemented",
		"method":  "PATCH",
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

// Logout handles user logout
func Logout(c *gin.Context) {
	// TODO: Implement logout logic (invalidate token)
	response.Success(c, gin.H{
		"message": "Logout functionality to be implemented",
	})
}

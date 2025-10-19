package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/komkem01/easy-attend-service/controller/auth"
	"github.com/komkem01/easy-attend-service/middlewares"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes() *gin.Engine {
	// Set gin mode based on environment
	gin.SetMode(gin.ReleaseMode) // Change to gin.DebugMode for development

	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "Easy Attend Service is running",
		})
	})

	// API version 1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes (public)
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/login", auth.Login)
			authRoutes.POST("/register", auth.Register)
			authRoutes.POST("/refresh", auth.RefreshToken)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			// User profile routes
			protected.GET("/profile", auth.GetProfile)
			protected.PATCH("/profile", auth.UpdateProfile) // Partial update
			protected.PUT("/profile", auth.ReplaceProfile)  // Full replacement
			protected.POST("/logout", auth.Logout)

			// Add more protected routes here as you develop features
		}
	}

	return router
}

// corsMiddleware handles CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

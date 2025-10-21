package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/komkem01/easy-attend-service/controller/auth"
	"github.com/komkem01/easy-attend-service/middlewares"
	"github.com/uptrace/bun"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(db *bun.DB) *gin.Engine {
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

	// Initialize services
	classroomService := auth.NewClassroomService(db)

	// Initialize controllers
	classroomController := auth.NewClassroomController(classroomService)
	assignmentController := auth.NewAssignmentController(db)

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

		// Public routes (no authentication required)
		public := v1.Group("/")
		{
			// System information
			public.GET("/info", auth.GetInfo)

			// View any user's profile by ID (public information only)
			public.GET("/profile/:id", auth.GetProfileByID)

			// Reference data endpoints
			public.GET("/genders", auth.GetGenders)
			public.GET("/genders/:id", auth.GetGendersByID)
			public.GET("/prefixes", auth.GetPrefixes)
			public.GET("/prefixes/:id", auth.GetPrefixesByID)

			// Schools (public read access)
			public.GET("/schools", auth.GetSchools)
			public.GET("/schools/:id", auth.GetSchoolByID)

			// Classrooms (public read access)
			public.GET("/classrooms", classroomController.GetClassrooms)
			public.GET("/classrooms/:id", classroomController.GetClassroom)

			// Assignments (public read access)
			public.GET("/assignments", assignmentController.GetAssignments)
			public.GET("/assignments/:id", assignmentController.GetAssignment)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			// User profile routes
			protected.GET("/profile", auth.GetProfile)
			protected.PATCH("/profile", auth.UpdateProfile)
			protected.PUT("/profile", auth.ReplaceProfile)
			protected.POST("/logout", auth.Logout)

			// Schools management (protected - requires authentication)
			protected.POST("/schools", auth.CreateSchool)
			protected.PATCH("/schools/:id", auth.UpdateSchool)
			protected.DELETE("/schools/:id", auth.DeleteSchool)

			// Classrooms management (protected - requires authentication)
			protected.POST("/classrooms", classroomController.CreateClassroom)
			protected.PATCH("/classrooms/:id", classroomController.UpdateClassroom)
			protected.DELETE("/classrooms/:id", classroomController.DeleteClassroom)

			// Assignments management (protected - requires authentication)
			protected.POST("/assignments", assignmentController.CreateAssignment)
			protected.PATCH("/assignments/:id", assignmentController.UpdateAssignment)
			protected.DELETE("/assignments/:id", assignmentController.DeleteAssignment)
			protected.POST("/assignments/:id/publish", assignmentController.PublishAssignment)

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

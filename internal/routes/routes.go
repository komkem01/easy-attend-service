package routes

import (
	"easy-attend-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Health check routes
		v1.GET("/health", handlers.HealthCheck)
		v1.GET("/version", handlers.GetVersion)

		// Future routes will be added here
		// attendance := v1.Group("/attendance")
		// users := v1.Group("/users")
		// classes := v1.Group("/classes")
	}
}

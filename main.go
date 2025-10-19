package main

import (
	"flag"
	"log"
	"net/http"

	"easy-attend-service/internal/config"
	"easy-attend-service/internal/database"
	"easy-attend-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Parse command line flags
	migrate := flag.Bool("migrate", false, "Run database migration only")
	force := flag.Bool("force", false, "Force migration even if there are errors")
	flag.Parse()

	log.Println("Starting Easy Attend Service...")

	// Load configuration
	cfg := config.Load()
	log.Println("Configuration loaded successfully")

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// If migrate flag is set, run migration and exit
	if *migrate {
		log.Println("Running database migration...")
		if err := database.Migrate(); err != nil {
			if *force {
				log.Printf("Migration failed but continuing due to --force flag: %v", err)
			} else {
				log.Printf("Database migration failed: %v", err)
				log.Println("If tables already exist, this is normal.")
				log.Println("Use --force flag to ignore migration errors")
				log.Fatal("Migration failed")
			}
		} else {
			log.Println("Database migration completed successfully!")
		}
		return
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	port := ":" + cfg.Server.Port
	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(port))
}

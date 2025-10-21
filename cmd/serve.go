package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "github.com/komkem01/easy-attend-service/configs"
	"github.com/komkem01/easy-attend-service/routes"
	"github.com/spf13/cobra"
)

// Serve command starts the HTTP server
func Serve() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		Long:  "Start the HTTP server for the Easy Attend Service API",
		Args:  NotReqArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return config.Open(cmd.Context())
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return config.Close(cmd.Context())
		},
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
	return cmd
}

func startServer() {
	// Get database connection
	db := config.Database()

	// Initialize routes with database
	router := routes.SetupRoutes(db)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("Server starting on port %s...\n", port)
	log.Printf("Server is running on http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/komkem01/easy-attend-service/cmd"
	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	// Root command
	rootCmd := &cobra.Command{
		Use:   "easy-attend-service",
		Short: "Easy Attend Service - Backend API",
		Long:  "A backend service for attendance management system",
	}

	// Add migrate commands
	rootCmd.AddCommand(cmd.Migrate())

	// Add serve command for HTTP server
	rootCmd.AddCommand(cmd.Serve())

	// Add healthcheck command
	rootCmd.AddCommand(cmd.Healthcheck())

	// Add version command
	rootCmd.AddCommand(cmd.VersionCmd())

	// Execute root command
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

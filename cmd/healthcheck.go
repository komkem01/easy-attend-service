package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Healthcheck command checks if the server is running
func Healthcheck() *cobra.Command {
	var host string
	var port string

	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "Check if the server is healthy",
		Long:  "Check if the Easy Attend Service server is running and responding",
		Args:  NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			performHealthcheck(host, port)
		},
	}

	cmd.Flags().StringVarP(&host, "host", "H", "localhost", "Server host to check")
	cmd.Flags().StringVarP(&port, "port", "p", "", "Server port to check (default from PORT env or 8080)")

	return cmd
}

func performHealthcheck(host, port string) {
	// Get port from flag, environment, or use default
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
	}

	url := fmt.Sprintf("http://%s:%s/health", host, port)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("🔍 Checking server health at %s...\n", url)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ Health check failed: %v\n", err)
		fmt.Println("💡 Make sure the server is running with: go run main.go serve")
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("✅ Server is healthy! Status: %s\n", resp.Status)
		fmt.Printf("🚀 Server is running on port %s\n", port)
		os.Exit(0)
	} else {
		fmt.Printf("⚠️  Server responded but with status: %s\n", resp.Status)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "1.0.0"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

// Version command shows version information
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Show version information for Easy Attend Service",
		Args:  NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			showVersion()
		},
	}
	return cmd
}

func showVersion() {
	fmt.Printf("Easy Attend Service\n")
	fmt.Printf("Version:      %s\n", Version)
	fmt.Printf("Go version:   %s\n", runtime.Version())
	fmt.Printf("OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Build time:   %s\n", BuildTime)
	fmt.Printf("Git commit:   %s\n", GitCommit)
}

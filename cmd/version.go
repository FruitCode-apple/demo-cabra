package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of demo-cabra",
	Long:  `version prints the build version, Go runtime version, and OS/architecture.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("demo-cabra version : %s\n", Version)
		fmt.Printf("Go version         : %s\n", runtime.Version())
		fmt.Printf("OS/Arch            : %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

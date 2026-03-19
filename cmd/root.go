package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "demo-cabra",
	Short: "A Cobra CLI demo for Debian development",
	Long: `demo-cabra is a small demonstration CLI application built with
the github.com/spf13/cobra library (packaged in Debian as
golang-github-spf13-cobra-dev).

It showcases how to structure a Go CLI project that is ready for
packaging in a Debian/Ubuntu environment.`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(versionCmd)
}

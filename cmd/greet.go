package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	greetName  string
	greetLoud  bool
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Print a greeting message",
	Long: `greet prints a personalised greeting.

Use --name to specify who to greet, and --loud to shout the message.`,
	Example: `  demo-cabra greet
  demo-cabra greet --name Alice
  demo-cabra greet --name Bob --loud`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := fmt.Sprintf("Hello, %s!", greetName)
		if greetLoud {
			msg = fmt.Sprintf("HELLO, %s!!!", greetName)
		}
		fmt.Println(msg)
	},
}

func init() {
	greetCmd.Flags().StringVarP(&greetName, "name", "n", "World", "Name of the person to greet")
	greetCmd.Flags().BoolVarP(&greetLoud, "loud", "l", false, "Shout the greeting in uppercase")
}

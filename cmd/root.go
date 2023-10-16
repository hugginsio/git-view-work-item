package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.huggins.io/prj/git-vwi/lib/git"
)

var rootCmd = &cobra.Command{
	Use:   "git-vwi",
	Short: "Git add-on for opening work item details in your browser based on the current branch.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, _ []string) {
		if cmd.Use != "version" {
			// NOTE: Global preflight. Runs before any commands are executed,
			// excluding built-in help messages.
			git.Version()
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("git-vwi")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

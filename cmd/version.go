package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.huggins.io/prj/git-vwi/lib/git"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information about git-vwi",
	Long: `Display version information about git-vwi.

The version of git-vwi and the available version of git is printed to
the standard output.`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("git-vwi version WIP")
		fmt.Println(git.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

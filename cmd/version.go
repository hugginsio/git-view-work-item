package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.huggins.io/prj/git-vwi/lib/git"
)

var (
	Version    = "1.5.1" // x-release-please-version
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display version information about git-vwi",
		Long: `Display version information about git-vwi.

The version of git-vwi and the available version of git is printed to
the standard output.`,
		Args: cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("git-vwi version %s\n", Version)
			fmt.Println(git.Version())
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

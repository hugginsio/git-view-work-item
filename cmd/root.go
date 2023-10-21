package cmd

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"text/template"

	"github.com/spf13/cobra"
	"go.huggins.io/prj/git-vwi/lib/git"
	"go.huggins.io/prj/git-vwi/lib/shell"
	"go.huggins.io/prj/git-vwi/lib/util"
)

var rootCmd = &cobra.Command{
	Use:   "git-vwi",
	Short: "Git add-on for opening work item details in your browser based on the current branch.",
	Long: `Git add-on for opening work item details in your browser based on the
current branch. This add-on relies on configuration through Git's global properties:

[git-view-work-item]
    url = "https://dev.azure.com/org/project/_workitems/edit/{{ .Identifier }}"
    regex = "[0-9]+"

You can use Go text templates to insert the following properties into the URL:
- Identifier (the identifier extracted from the current branch name)

Learn more about Go text templates at https://pkg.go.dev/text/template`,
	PreRun: func(cmd *cobra.Command, _ []string) {
		if cmd.Use != "version" {
			// NOTE: Global preflight. Runs before any commands are executed, excluding built-in help messages.
			git.Version()
			git.RepositoryCheck()
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		urlConfig := git.GetConfig("git-view-work-item.url")
		regexConfig := git.GetConfig("git-view-work-item.regex")

		if len(urlConfig) == 0 {
			fmt.Println("fatal: URL configuration property not found")
			fmt.Println("run `git-vwi -h` or `git-vwi help` flag to view help")
			os.Exit(4)
		} else if len(regexConfig) == 0 {
			fmt.Println("fatal: identifier regular expression configuration property not found")
			fmt.Println("run `git-vwi -h` or `git-vwi help` flag to view help")
			os.Exit(4)
		}

		currentBranch := git.CurrentBranch()
		regex, err := regexp.Compile(regexConfig)
		util.CheckErrorFatal(err, "regular expression failed to compile")

		type Properties struct {
			Identifier string
			// TODO: add more properties, like the repository name or current directory
		}

		if !regex.MatchString(currentBranch) {
			fmt.Println("fatal: no matches for regular expression")
			fmt.Println("")
			fmt.Println("Current git branch: " + currentBranch)
			fmt.Println("Regular expression: " + regex.String())
			os.Exit(5)
		}

		var props Properties
		props.Identifier = regex.FindString(currentBranch)

		var finalUrl bytes.Buffer
		templ := template.Must(template.New("templ").Parse(urlConfig))
		templ.Execute(&finalUrl, props)

		fmt.Println(finalUrl.String())
		// TODO: make compatible with non-macOS environments
		shell.Execute("open", finalUrl.String())
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
	// TODO: provide flag that will copy the link instead of opening it in the browser
}

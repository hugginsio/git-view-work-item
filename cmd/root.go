package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"go.huggins.io/prj/git-vwi/lib/git"
	"go.huggins.io/prj/git-vwi/lib/shell"
	"go.huggins.io/prj/git-vwi/lib/util"
)

var (
	rootCmdBoolCopy bool
	rootCmd         = &cobra.Command{
		Use:   "git-vwi",
		Short: "Git add-on for opening work item details in your browser based on the current branch.",
		Long: `Git add-on for opening work item details in your browser based on the
current branch. This add-on relies on configuration through Git's global properties:

[git-view-work-item]
    url = "https://dev.azure.com/org/project/_workitems/edit/{{ .Identifier }}"
    regex = "[0-9]+"

You can use Go text templates to insert the following properties into the URL:

- Identifier (the identifier extracted from the current branch name)
- Directory: the current directory name (but not the full path).
- Identifier: the identifier extracted from the current branch name.
- Repository: the repository name, taken from remote.origin.url.
- Url: the URL of the repository, taken from remote.origin.url.

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

			// TODO: extract repository name with (?!.*\/).+[^\.git]

			type Properties struct {
				Directory  string
				Identifier string
				Repository string
				Url        string
			}

			var props Properties

			cwd, _ := os.Getwd()
			props.Directory = filepath.Base(cwd)

			currentBranch := git.CurrentBranch()
			branchRegex, err := regexp.Compile(regexConfig)
			util.CheckErrorFatal(err, "regular expression failed to compile")

			if !branchRegex.MatchString(currentBranch) {
				fmt.Println("fatal: no matches for regular expression")
				fmt.Println("")
				fmt.Println("Current git branch: " + currentBranch)
				fmt.Println("Regular expression: " + branchRegex.String())
				os.Exit(5)
			}

			props.Identifier = branchRegex.FindString(currentBranch)

			remoteUrl := git.GetConfig("remote.origin.url")
			remoteUrl = strings.ReplaceAll(remoteUrl, ".git", "")

			urlParts := strings.Split(remoteUrl, "/")
			props.Repository = urlParts[len(urlParts)-1]
			props.Url = remoteUrl

			var finalUrl bytes.Buffer
			templ := template.Must(template.New("templ").Parse(urlConfig))
			templ.Execute(&finalUrl, props)

			fmt.Println(finalUrl.String())
			// TODO: make compatible with non-macOS environments
			if !rootCmdBoolCopy {
				shell.Execute("open", finalUrl.String())
			} else {
				clipboard.WriteAll(finalUrl.String())
			}
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&rootCmdBoolCopy, "copy", "c", rootCmdBoolCopy, "copy URL to clipboard instead of opening in browser")
}

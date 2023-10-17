package git

import (
	"regexp"
	"strings"

	"go.huggins.io/prj/git-vwi/lib/shell"
	"go.huggins.io/prj/git-vwi/lib/util"
)

var (
	git404         string = "could not execute git"
	gitConfigUrl   string = "git-view-work-item.url"
	gitConfigRegex string = "git-view-work-item.regex"
)

// Returns the version of git currently in your path.
func Version() string {
	out, err := shell.Execute("git", "version")
	util.CheckErrorFatal(err, git404)

	versionOutput := strings.TrimSpace(string(out))
	versionRegex := regexp.MustCompile(`[0-9].*`)
	versionString := versionRegex.FindString(versionOutput)

	return versionString
}

// No action if you are inside a repository. If you are not in a repository,
// display a message to the user and exit the program.
func RepositoryCheck() {
	_, err := shell.Execute("git", "rev-parse", "--is-inside-work-tree")
	util.CheckErrorFatal(err, "not a git repository (or any of the parent directories)")
}

// Returns the name of the current branch.
func CurrentBranch() string {
	out, err := shell.Execute("git", "branch", "--show-current")
	util.CheckErrorFatal(err, git404)
	return strings.TrimSpace(string(out))
}

// Get the value of a global config key.
func GetConfig(key string) string {
	// NOTE: potential errors ignored, since PersistentPreRun will have validated
	// that Git is executable and we are in a repository.
	out, err := shell.Execute("git", "config", "--get", key)

	if err != nil {
		return ""
	} else {
		return strings.TrimSpace(string(out))
	}
}

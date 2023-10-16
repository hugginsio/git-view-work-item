package git

import (
	"regexp"
	"strings"

	"go.huggins.io/prj/git-vwi/lib/shell"
	"go.huggins.io/prj/git-vwi/lib/util"
)

var git404 string = "could not execute git"

// Returns the version of git currently in your path.
func Version() string {
	out, err := shell.Execute("git", "version")
	util.CheckErrorFatal(err, git404)

	versionOutput := strings.TrimSpace(string(out))
	versionRegex := regexp.MustCompile(`[0-9].*`)
	versionString := versionRegex.FindString(versionOutput)

	return versionString
}

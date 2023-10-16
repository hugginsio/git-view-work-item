package shell

import (
	"os/exec"
)

// Execute wraps `exec.Command` and provide the Output(). Any returned error
// will usually be of type \*ExitError. If c.Stderr was nil, Output
// populates ExitError.Stderr.
func Execute(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}

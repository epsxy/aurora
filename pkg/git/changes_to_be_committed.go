package git

import (
	"strings"
)

// AreChangesAddedToBeCommited checks if `git status` output contains
// `Changes to be committed` string, meaning there are currently changes
// staged for commit.
func AreChangesAddedToBeCommited() bool {
	status := GitStatus()

	if strings.Contains(status, "Changes to be committed") {
		return true
	}
	return false
}

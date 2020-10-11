package git

import (
	"log"
	"os/exec"
	"strings"
)

// Status runs `git status` command.
// Exits if it fails.
func Status() string {
	cmd := exec.Command("git", "status")
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

// AreChangesAddedToBeCommited checks if `git status` output contains
// `Changes to be committed` string, meaning there are currently changes
// staged for commit.
func AreChangesAddedToBeCommited() bool {
	status := Status()

	if strings.Contains(status, "Changes to be committed") {
		return true
	}
	return false
}

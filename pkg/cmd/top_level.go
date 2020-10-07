package cmd

import (
	"log"
	"os/exec"
)

// ShowTopLevel runs `git rev-parse --show-toplevel` command.
// Exits if it fails.
func ShowTopLevel() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()

	if err != nil {
		log.Fatal("Unable to retrieve repository path.\nAre you in a git repository?\nAborting...")
	}
	return string(out)
}

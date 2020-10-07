package cmd

import (
	"log"
	"os/exec"
)

// GitStatus runs `git status` command.
// Exits if it fails.
func GitStatus() string {
	cmd := exec.Command("git", "status")
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

package cmd

import (
	"log"
	"os/exec"
)

// GitLog runs `git log` command.
// Exits if it fails.
func GitLog() string {
	cmd := exec.Command("git", "log", "--abbrev-commit")

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

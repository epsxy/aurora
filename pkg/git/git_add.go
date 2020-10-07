package git

import (
	"log"
	"os/exec"
)

// GitAddAll runs `git add -A` command.
// Exits if it fails.
func GitAddAll() {
	cmd := exec.Command("git", "add", "-A")
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

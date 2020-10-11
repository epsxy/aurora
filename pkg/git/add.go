package git

import (
	"log"
	"os/exec"
)

// AddAll runs `git add -A` command.
// Exits if it fails.
func AddAll() {
	cmd := exec.Command("git", "add", "-A")
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

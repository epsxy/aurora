package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

// GitCommit runs `git commit -m` command.
// Exits if it fails.
func GitCommit(m string) {
	cmd := exec.Command("git", "commit", "-m", m)
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}

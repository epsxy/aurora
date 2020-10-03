package cmd

import (
	"fmt"
	"os/exec"
)

// GitCommit : Exec git commit command
func GitCommit(m string) {
	cmd := exec.Command("git", "commit", "-m", m)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", out)
}

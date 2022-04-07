package git

import (
	"log"
	"os/exec"
)

func TagCommit(version string) {
	cmd := exec.Command("git", "tag", version)
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

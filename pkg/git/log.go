package git

import (
	"fmt"
	"log"
	"os/exec"
)

// Log runs `git log --abrev-commit` command.
// Exits if it fails.
func Log() string {
	cmd := exec.Command("git", "log", "--abbrev-commit")

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

//LogChangelog runs `git log --oneline --first-parent [rev1]...[rev2]`.
//It logs the commit history between 2 revisions.
//If the first one is empty, defaults to `HEAD`
func LogChangelog(firstCommit string, secondCommit string) string {
	if firstCommit == "" {
		firstCommit = "HEAD"
	}
	opts := fmt.Sprintf("%s...%s", firstCommit, secondCommit)
	cmd := exec.Command("git", "log", "--oneline", "--first-parent", opts)

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

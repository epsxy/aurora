package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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

//LogShortHash returns commit short hashes list between 2 revisions
func LogShortHash(rev1 string, rev2 string) []string {
	if rev1 == "" {
		rev1 = "HEAD"
	}
	opts := fmt.Sprintf("%s...%s", rev1, rev2)
	cmd := exec.Command("git", "log", "--oneline", "--first-parent", "--pretty=format:%h", opts)

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(out), "\n")
}

//LogCommitMessage returns commit message list between 2 revisions
func LogCommitMessage(rev1 string, rev2 string) []string {
	if rev1 == "" {
		rev1 = "HEAD"
	}
	opts := fmt.Sprintf("%s...%s", rev1, rev2)
	cmd := exec.Command("git", "log", "--oneline", "--first-parent", "--pretty=format:%s", opts)

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	if len(out) == 0 {
		return []string{}
	}
	return strings.Split(string(out), "\n")
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

//ParseLogOutput generate a list of git.Commit from a git log output
func ParseLogOutput(log string) []Commit {
	var c *Commit
	logSplit := strings.Split(log, "\n")
	res := make([]Commit, len(logSplit)-1)
	for i, line := range logSplit[:len(logSplit)-1] {
		split := strings.Split(line, " ")
		h := split[0]
		c = &Commit{
			Hash: h,
		}
		m := strings.Join(split[1:], " ")
		c.FillFromStrMsg(m)
		res[i] = *c
	}
	return res
}

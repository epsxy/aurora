package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//Commit struct respresents a git commit object
type Commit struct {
	Hash             string
	Type             string
	Scope            string
	IsBreakingChange bool
	IsMergeCommit    bool
	Msg              string
}

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

//FillFromStrMsg fill commit information from a string inside a git.Commit object
func (c *Commit) FillFromStrMsg(m string) {
	cType := ""
	scope := ""
	isBreakingChange := false
	msg := ""

	// FIXME: Dirty fix to identify default merge commits
	if strings.Contains(m, "Merge pull request") {
		(*c).Type = ""
		(*c).Scope = ""
		(*c).IsBreakingChange = false
		(*c).IsMergeCommit = true
		(*c).Msg = msg
		return
	}

	// If `:` is missing, this commit does not respects conventional commits.
	if !strings.Contains(m, ":") {
		(*c).Type = "misc"
		(*c).Scope = ""
		(*c).IsBreakingChange = false
		(*c).IsMergeCommit = false
		(*c).Msg = m
		return
	}

	globalSplit := strings.Split(m, ":")
	funcPart := globalSplit[0]
	msg = strings.Join(globalSplit[1:], "")
	if strings.Contains(funcPart, "!") {
		funcPart = strings.ReplaceAll(funcPart, "!", "")
		isBreakingChange = true
	}
	cType = funcPart
	if strings.Contains(funcPart, "(") {
		scopePart := strings.Split(funcPart, "(")
		cType = scopePart[0]
		scope = strings.ReplaceAll(scopePart[1], ")", "")
	}

	(*c).Type = cType
	(*c).Scope = scope
	(*c).IsBreakingChange = isBreakingChange
	(*c).IsMergeCommit = false
	(*c).Msg = msg
}

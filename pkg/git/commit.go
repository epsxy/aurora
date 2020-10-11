package git

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
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

// CreateCommit runs `git commit -m` command.
// Exits if it fails.
func CreateCommit(m string) {
	cmd := exec.Command("git", "commit", "-m", m)
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}

//LintCommitMessage check if a message respect the conventional commit convention
// TODO: apply .env file
func LintCommitMessage(m string) bool {
	r := "^(revert: )?(feat|fix|docs|style|refactor|perf|test|chore|ci)(\\(.+\\))?(\\!)?: .{1,72}"
	isMatching, _ := regexp.MatchString(r, m)
	return isMatching
}

//IsMergeCommit checks if a commit is a merge commit
// FIXME: Dirty fix to identify default merge commits
func IsMergeCommit(m string) bool {
	return strings.Contains(m, "Merge pull request")
}

//FillFromStrMsg fill commit information from a string inside a git.Commit object
func (c *Commit) FillFromStrMsg(m string) {
	cType := ""
	scope := ""
	isBreakingChange := false
	msg := ""

	if IsMergeCommit(m) {
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

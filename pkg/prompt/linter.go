package prompt

import (
	"fmt"
	"log"

	"github.com/epsxy/gommitizen/pkg/git"
)

//Lint displays linter prompt
func Lint(firstCommit string, secondCommit string) {
	fails := false
	messages := git.LogCommitMessage(firstCommit, secondCommit)
	for _, m := range messages {
		if m != "" && !git.LintCommitMessage(m) {
			if git.IsMergeCommit(m) {
				continue
			}
			fails = true
			fmt.Println("Invalid commit:")
			fmt.Println(m)
		}
	}
	if fails {
		log.Fatal("Some commits are not valid")
	}
}

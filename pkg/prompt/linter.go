package prompt

import (
	"os"

	"github.com/epsxy/gommitizen/pkg/git"
	"github.com/gookit/color"
)

//Lint displays linter prompt
func Lint(firstCommit string, secondCommit string) {
	fails := false
	messages := git.LogCommitMessage(firstCommit, secondCommit)
	successSymbol := "✔"
	failSymbol := "✘"
	for _, m := range messages {
		if m != "" && !git.LintCommitMessage(m) {
			if git.IsMergeCommit(m) {
				color.Cyan.Printf("%s  %s\n", successSymbol, m)
				continue
			}
			fails = true
			color.Red.Printf("%s  %s\n", failSymbol, m)
			continue
		}
		color.Cyan.Printf("%s  %s\n", successSymbol, m)
	}
	if fails {
		color.Red.Printf("Linting failed, some commits are not valid\n")
		os.Exit(1)
	}
}

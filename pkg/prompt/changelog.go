package prompt

import (
	"fmt"

	"github.com/epsxy/aurora/pkg/git"
	"github.com/epsxy/aurora/pkg/parser"
)

//Changelog displays changelog generator prompt
func Changelog(firstCommit string, secondCommit string) {
	log := git.LogChangelog(firstCommit, secondCommit)
	commits := git.ParseLogOutput(log)
	md := parser.GenerateMarkdownChangelog(commits)
	fmt.Println(md)
}

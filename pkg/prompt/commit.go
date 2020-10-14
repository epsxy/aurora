package prompt

import (
	"fmt"

	"github.com/epsxy/gommitizen/pkg/cli"
	"github.com/epsxy/gommitizen/pkg/git"
	"github.com/epsxy/gommitizen/pkg/parser"
)

type commitsConf struct {
	Types  []string `yaml:"types,flow"`
	Scopes []string `yaml:"scopes,flow"`
}

// Commit : gommitizen commit prompt
func Commit() string {

	if git.AreChangesAddedToBeCommited() == false {
		fmt.Println("No file staged for commit")
		cli.StageAll()
	}

	conf := parser.EnvFileParser()

	t := cli.CommitType(conf.Types)
	s := cli.CommitScope(conf.Scopes)
	b := cli.BreakingChange()
	m := cli.CommitShortMsg()

	if s == "" {
		return fmt.Sprintf("%s%s: %s", t, b, m)
	}

	return fmt.Sprintf("%s(%s)%s: %s", t, s, b, m)
}

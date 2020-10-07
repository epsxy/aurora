package cli

import (
	"fmt"

	"github.com/epsxy/gommitizen/pkg/git"

	"github.com/epsxy/gommitizen/pkg/parser"
)

type commitsConf struct {
	Types  []string `yaml:"types,flow"`
	Scopes []string `yaml:"scopes,flow"`
}

// GitPrompt : Global entrypoint
func GitPrompt() string {

	if git.AreChangesAddedToBeCommited() == false {
		fmt.Println("No file staged for commit")
		stageAll()
	}

	conf := parser.EnvFileParser()

	t := commitType(conf.Types)
	s := commitScope(conf.Scopes)
	b := breakingChange()
	m := commitShortMsg()

	if s == "" {
		return fmt.Sprintf("%s%s: %s", t, b, m)
	}

	return fmt.Sprintf("%s(%s)%s: %s", t, s, b, m)
}

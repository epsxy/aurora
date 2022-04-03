package prompt

import (
	"fmt"
	"log"

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
		stageAll := cli.YesNoSelect("Stage all files", false)
		if stageAll == true {
			git.AddAll()
		} else {
			log.Fatal("Staging no file. Aborted.")
		}
	}

	conf := parser.GetConf()

	types := conf.Types
	if types == nil {
		types = []string{
			"feat",
			"fix",
			"build",
			"chore",
			"ci",
			"docs",
			"style",
			"refactor",
			"perf",
			"test",
		}
	}
	scopes := conf.Scopes

	// Type
	t := cli.Select("Commit type", types)

	// Scope
	var s string
	if scopes == nil {
		s = cli.MessageInput("Commit scope", -1)
	} else {
		s = cli.Select("Commit scope", scopes)
	}

	// Breaking change
	b := ""
	if cli.YesNoSelect("Breaking change", true) == true {
		b = "!"
	}

	// Short message
	m := cli.MessageInput("Commit short message", 72)

	if s == "" {
		return fmt.Sprintf("%s%s: %s", t, b, m)
	}
	return fmt.Sprintf("%s(%s)%s: %s", t, s, b, m)
}

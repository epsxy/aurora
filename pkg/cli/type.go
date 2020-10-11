package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

func commitType(conf []string) string {
	label := "Commit type"

	if conf == nil {
		conf = []string{
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

	prompt := promptui.Select{
		Label:  label,
		Items:  conf,
		Stdout: &BellSkipper{},
	}
	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

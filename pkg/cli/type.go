package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

func commitType() string {
	label := "Commit type"
	types := []string{"feat", "fix", "build", "chore", "ci", "docs", "style", "refactor", "perf", "test"}

	prompt := promptui.Select{
		Label: label,
		Items: types,
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

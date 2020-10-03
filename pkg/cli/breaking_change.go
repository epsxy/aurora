package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

func breakingChange() string {
	label := "BreakingChange"
	types := []string{"No", "Yes"}

	prompt := promptui.Select{
		Label: label,
		Items: types,
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	if res == "Yes" {
		return "!"
	}
	return ""
}

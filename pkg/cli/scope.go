package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

func commitScope() string {
	label := "Commit scope"

	prompt := promptui.Prompt{
		Label: label,
	}

	res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

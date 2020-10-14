package cli

import (
	"log"

	"github.com/epsxy/gommitizen/pkg/git"
	"github.com/manifoldco/promptui"
)

// StageAll displays a prompt to ask if the user wants to add all files
func StageAll() {
	label := "Stage all files"
	types := []string{"Yes", "No"}

	prompt := promptui.Select{
		Label:  label,
		Items:  types,
		Stdout: &BellSkipper{},
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	if res == "Yes" {
		git.AddAll()
	} else {
		log.Fatal("Staging no file. Aborted.")
	}
}

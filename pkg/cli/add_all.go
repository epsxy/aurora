package cli

import (
	"log"

	"github.com/epsxy/gommitizen/pkg/cmd"

	"github.com/manifoldco/promptui"
)

func stageAll() {
	label := "Stage all files"
	types := []string{"Yes", "No"}

	prompt := promptui.Select{
		Label: label,
		Items: types,
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	if res == "Yes" {
		cmd.GitAddAll()
	} else {
		log.Fatal("Staging no file. Aborted.")
	}
}

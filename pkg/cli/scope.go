package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

// CommitScope displays a prompt to select commit scope.
// If scopes are prefined in the .gommitizen.yml conf file,
// this input is replaced by a select component.
func CommitScope(conf []string) string {
	label := "Commit scope"

	if conf != nil {
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

	prompt := promptui.Prompt{
		Label:  label,
		Stdout: &BellSkipper{},
	}

	res, err := prompt.Run()
	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

// BreakingChange displays a prompt to ask if current commit
// introduces a breaking change or not
func BreakingChange() string {
	label := "BreakingChange"
	types := []string{"No", "Yes"}

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
		return "!"
	}
	return ""
}

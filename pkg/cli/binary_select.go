package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

// YesNoSelect display a select promptui item
// Yes/No if negative = false, No/Yes if negative = true
func YesNoSelect(label string, negative bool) bool {
	var items []string
	if negative == true {
		items = []string{"No", "Yes"}
	} else {
		items = []string{"Yes", "No"}
	}
	prompt := promptui.Select{
		Label:  label,
		Items:  items,
		Stdout: &BellSkipper{},
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	if res == "Yes" {
		return true
	}
	return false
}

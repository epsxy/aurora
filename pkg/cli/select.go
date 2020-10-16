package cli

import (
	"log"

	"github.com/manifoldco/promptui"
)

// Select display a select promptui item
func Select(label string, items []string) string {
	prompt := promptui.Select{
		Label:  label,
		Items:  items,
		Stdout: &BellSkipper{},
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}
	return res
}

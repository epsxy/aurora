package cli

import (
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

// Select display a select promptui item
func Select(label string, items []string) string {
	searcher := func(input string, index int) bool {
		t := items[index]
		name := strings.Replace(strings.ToLower(t), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		Stdout:   &BellSkipper{},
		Searcher: searcher,
	}

	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}
	return res
}

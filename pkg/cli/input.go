package cli

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func validate(max int) func(string) error {
	return func(input string) error {
		if max < 0 {
			return nil
		}
		if len(input) > max {
			return fmt.Errorf("String length should not be > %d", max)
		}
		return nil
	}
}

// MessageInput displays an input to type commit short message
func MessageInput(label string, max int) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate(max),
		Stdout:   &BellSkipper{},
	}

	res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

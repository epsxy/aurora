package cli

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
)

func validateMsgLength(input string) error {
	if len(input) > 72 {
		return errors.New("String length should not be > 72")
	}
	return nil
}

// CommitShortMsg displays an input to type commit short message
func CommitShortMsg() string {
	label := "Commit short message"

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateMsgLength,
		Stdout:   &BellSkipper{},
	}

	res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

package cli

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
)

func validateMsgLength(input string) error {
	if len(input) > 90 {
		return errors.New("String length should not be > 90")
	}
	return nil
}

func commitShortMsg() string {
	label := "Commit short message"

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateMsgLength,
	}

	res, err := prompt.Run()

	if err != nil {
		log.Fatal("Aborted")
	}

	return res
}

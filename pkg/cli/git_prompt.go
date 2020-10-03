package cli

import (
	"fmt"
)

// GitPrompt : Global entrypoint
func GitPrompt() string {

	t := commitType()
	s := commitScope()
	b := breakingChange()
	m := commitShortMsg()

	if s == "" {
		return fmt.Sprintf("%s%s: %s", t, b, m)
	}

	return fmt.Sprintf("%s(%s)%s: %s", t, s, b, m)
}

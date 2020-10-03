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

	return fmt.Sprintf("%s(%s)%s: %s", t, s, b, m)
}

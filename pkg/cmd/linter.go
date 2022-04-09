package cmd

import (
	"github.com/epsxy/aurora/pkg/prompt"
	"github.com/spf13/cobra"
)

//Lint is the lint command entrypoint
var Lint = &cobra.Command{
	Use:   "lint",
	Short: "Run commit linter",
	Long: `
Lint commits between git revisions.
Supported syntax for example:
	- aurora lint v1.0.0
	- aurora lint feat-branch-1 master
	- aurora lint 0e4bc1a 7e1abc8
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		SetGlobalFlags(cmd)

		f := ""
		s := ""
		if len(args) == 1 {
			s = args[0]
		}
		if len(args) == 2 {
			f = args[0]
			s = args[1]
		}
		prompt.Lint(f, s)
	},
}

package cmd

import (
	"github.com/epsxy/gommitizen/pkg/prompt"
	"github.com/spf13/cobra"
)

//Changelog is the changelog command entrypoint
var Changelog = &cobra.Command{
	Use:   "changelog",
	Short: "Run changelog generator",
	Long: `
Generate a changelog between git revisions.
Supported syntax for example:
	- gommitizen changelog v1.0.0
	- gommitizen changelog feat-branch-1 master
	- gommitizen changelog 0e4bc1a 7e1abc8
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		f := ""
		s := ""
		if len(args) == 1 {
			s = args[0]
		}
		if len(args) == 2 {
			f = args[0]
			s = args[1]
		}
		prompt.Changelog(f, s)
	},
}

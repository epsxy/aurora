package cmd

import (
	"fmt"
	"os"

	"github.com/epsxy/gommitizen/pkg/cli"
	"github.com/epsxy/gommitizen/pkg/git"
	"github.com/epsxy/gommitizen/pkg/prompt"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gommitizen",
	Short: "Gommitizen is a commit formatter utility",
	Long:  "A Go program to help you create conventional commits",
}

var commit = &cobra.Command{
	Use:   "commit",
	Short: "Run commit formatter",
	Long:  "Create a commit following conventional commits convention",
	Run: func(cmd *cobra.Command, args []string) {
		message := cli.GitPrompt()
		git.GitCommit(message)
	},
}

var changelog = &cobra.Command{
	Use:   "changelog",
	Short: "Run changelog generator",
	Long: `
Generate a changelog between now and a previous revision or branch.
Supported synthax for example:
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

// Execute is the root entrypoint of the Cobra CLI
func Execute() {
	root.AddCommand(commit, changelog)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

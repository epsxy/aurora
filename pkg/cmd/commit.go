package cmd

import (
	"github.com/epsxy/gommitizen/pkg/cli"
	"github.com/epsxy/gommitizen/pkg/git"
	"github.com/spf13/cobra"
)

//Commit is the commit command entrypoint
var Commit = &cobra.Command{
	Use:   "commit",
	Short: "Run commit formatter",
	Long:  "Create a commit following conventional commits convention",
	Run: func(cmd *cobra.Command, args []string) {
		message := cli.GitPrompt()
		git.CreateCommit(message)
	},
}

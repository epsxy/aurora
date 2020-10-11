package cmd

import (
	"fmt"
	"os"

	"github.com/epsxy/gommitizen/pkg/cli"
	"github.com/epsxy/gommitizen/pkg/git"
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

// Execute is the root entrypoint of the Cobra CLI
func Execute() {
	root.AddCommand(commit)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

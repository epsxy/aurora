package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "aurora",
	Short:   "Aurora is a commit formatter utility",
	Long:    "A Go program to help you create conventional commits",
	Version: "v0.2.0",
}

// Execute is the root entrypoint of the Cobra CLI
func Execute() {
	root.AddCommand(Commit, Changelog, Lint)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

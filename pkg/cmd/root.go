package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "aurora",
	Short:   "Aurora is a commit formatter utility",
	Long:    "A Go program to help you create conventional commits",
	Version: "v0.0.0", // will be filled in `Execute` entrypoint
}

// Execute is the root entrypoint of the Cobra CLI
func Execute() {
	// init version number
	version, err := os.ReadFile("VERSION")
	if err != nil {
		log.Fatal("unexpected error: impossible to find VERSION")
	}
	root.Version = string(version)

	// add commands and flags
	root.AddCommand(Commit, Changelog, Lint, Release)
	Release.AddCommand(Major, Minor, Fix)
	root.PersistentFlags().Bool("dry-run", false, "enable dry-run mode")
	root.PersistentFlags().Bool("verbose", false, "enable verbose mode")

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

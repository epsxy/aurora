package cmd

import (
	"fmt"
	"os"

	"github.com/epsxy/aurora/pkg/global"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "aurora",
	Short:   "Aurora is a commit formatter utility",
	Long:    "A Go program to help you create conventional commits",
	Version: "v0.0.0", // will be filled in `Execute` entrypoint
}

func SetGlobalFlags(cmd *cobra.Command) {
	v, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		v = false
	}
	global.SetVerbose(v)

	d, err := cmd.Flags().GetBool("dry-run")
	if err != nil {
		d = false
	}
	global.SetDryRun(d)
}

// Execute is the root entrypoint of the Cobra CLI
func Execute(version string) {
	// setup version
	root.Version = version
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

package cmd

import (
	"github.com/epsxy/aurora/pkg/git"
	"github.com/epsxy/aurora/pkg/global"
	"github.com/epsxy/aurora/pkg/prompt"
	"github.com/spf13/cobra"
)

//Commit is the commit command entrypoint
var Commit = &cobra.Command{
	Use:   "commit",
	Short: "Run commit formatter",
	Long:  "Create a commit following conventional commits convention",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			v = false
		}
		global.SetVerbose(v)

		message := prompt.Commit()
		git.CreateCommit(message)
	},
}

package main

import (
	"github.com/epsxy/gommitizen/pkg/cli"
	cmd "github.com/epsxy/gommitizen/pkg/git"
)

func main() {
	s := cli.GitPrompt()
	cmd.GitCommit(s)
}

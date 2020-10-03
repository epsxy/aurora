package main

import (
	"github.com/epsxy/gommitizen/pkg/cli"
	"github.com/epsxy/gommitizen/pkg/cmd"
)

func main() {
	s := cli.GitPrompt()
	cmd.GitCommit(s)
}

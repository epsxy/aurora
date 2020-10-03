package main

import (
	"github.com/epsxy/gomitizen/pkg/cli"
	"github.com/epsxy/gomitizen/pkg/cmd"
)

func main() {
	s := cli.GitPrompt()
	cmd.GitCommit(s)
}

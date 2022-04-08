package main

import _ "embed"

import (
	"github.com/epsxy/aurora/pkg/cmd"
)

//go:embed VERSION
var version string

func main() {
	cmd.Execute(version)
}

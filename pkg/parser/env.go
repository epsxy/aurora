package parser

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

//GommitizenConf : .yaml configuration file struct
type GommitizenConf struct {
	Types  []string `yaml:"types,flow"`
	Scopes []string `yaml:"scopes,flow"`
}

func parseEnvFilePath() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()

	if err != nil {
		log.Fatal("Unable to retrieve repository path.\nAre you in a git repository?\nAborting...")
	}
	return strings.Replace(string(out), "\n", "", 1)
}

//EnvFileParser : parse gommitizen env file
func EnvFileParser() GommitizenConf {
	var conf GommitizenConf
	path := parseEnvFilePath()
	s := filepath.Join(path, ".gommitizen.yml")

	yamlFile, err := ioutil.ReadFile(s)
	if err != nil {
		return GommitizenConf{Scopes: nil, Types: nil}
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Printf("Malformed gommitizen configuration file: %v", err)
		return GommitizenConf{Scopes: nil, Types: nil}
	}
	return conf
}

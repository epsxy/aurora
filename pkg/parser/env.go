package parser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/epsxy/gommitizen/pkg/git"
	"gopkg.in/yaml.v2"
)

//GommitizenConf : .yaml configuration file struct
type GommitizenConf struct {
	Types  []string `yaml:"types,flow"`
	Scopes []string `yaml:"scopes,flow"`
}

func parseEnvFilePath() string {
	path := git.ShowTopLevel()
	return strings.Replace(path, "\n", "", 1)
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

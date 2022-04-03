package parser

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/epsxy/aurora/pkg/git"
	"gopkg.in/yaml.v2"
)

//AuroraConf : .yaml configuration file struct
type AuroraConf struct {
	Types  []string `yaml:"types,flow"`
	Scopes []string `yaml:"scopes,flow"`
}

func parseEnvFilePath() string {
	path := git.ShowTopLevel()
	return filepath.Join(strings.Replace(path, "\n", "", 1), ".aurora.yml")
}

func ParseHomeFilePath() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(h, ".aurora.yml")
}

//EnvFileParser : parse aurora env files
func GetConf() *AuroraConf {
	log.Printf("parsing env file")

	p := parseEnvFilePath()
	c, err := readConfFromPath(p)
	if err != nil && c != nil {
		return c
	}

	p = ParseHomeFilePath()
	c, err = readConfFromPath(p)
	if err != nil && c != nil {
		return c
	}

	return &AuroraConf{Scopes: nil, Types: nil}
}

func readConfFromPath(path string) (*AuroraConf, error) {
	var conf *AuroraConf

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

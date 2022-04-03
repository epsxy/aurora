package parser

import (
	"io/ioutil"
	"log"
	"os"
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
	return filepath.Join(strings.Replace(path, "\n", "", 1), ".gommitizen.yml")
}

func ParseHomeFilePath() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(h, ".gommitizen.yml")
}

//EnvFileParser : parse gommitizen env files
func GetConf() *GommitizenConf {
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

	return &GommitizenConf{Scopes: nil, Types: nil}
}

func readConfFromPath(path string) (*GommitizenConf, error) {
	var conf *GommitizenConf

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

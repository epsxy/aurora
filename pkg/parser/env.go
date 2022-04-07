package parser

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/epsxy/aurora/pkg/git"
	"github.com/epsxy/aurora/pkg/global"
	"gopkg.in/yaml.v2"
)

//AuroraConf : .yaml configuration file struct
type AuroraConf struct {
	Types  []string `yaml:"types,flow"`
	Scopes []string `yaml:"scopes,flow"`
}

func parseEnvFilePath() string {
	path := git.ShowTopLevel()
	return filepath.Join(path, ".aurora.yml")
}

func parseHomeFilePath() string {
	h, err := os.UserHomeDir()
	if err != nil {
		if global.GetVerbose() {
			log.Println("error while getting home dir")
		}
		return ""
	}
	return filepath.Join(h, ".aurora.yml")
}

//EnvFileParser : parse aurora env files
func GetConf() *AuroraConf {
	if global.GetVerbose() {
		log.Printf("parsing env file")
	}

	p := parseEnvFilePath()
	c, err := readConfFromPath(p)
	if err != nil && c != nil {
		return c
	}
	if global.GetVerbose() {
		log.Println("no conf file was found in repo, defaulting to ~/")
	}

	p = parseHomeFilePath()
	c, err = readConfFromPath(p)
	if err != nil && c != nil {
		return c
	}
	if global.GetVerbose() {
		log.Println("no conf file was found in home, defaulting to empty")
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

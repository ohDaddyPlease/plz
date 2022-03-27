package parser

import (
	"github.com/ohDaddyPlease/plz/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	UserHomeDir    string
	ConfigFilePath string
)

const ConfigFileName = ".plz"

func init() {
	var err error
	UserHomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	ConfigFilePath = filepath.Join(UserHomeDir, ConfigFileName)
}

func ParseArgs() model.CMD {
	var (
		args    = make([]string, 0)
		command string
	)

	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	c := model.CMD{
		Command: command,
		Args:    args,
	}
	return c
}

func ParseConfigFile() (model.ConfigurationFile, error) {
	content, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return model.ConfigurationFile{}, err
	}
	configurationFile := model.ConfigurationFile{}
	err = yaml.Unmarshal(content, &configurationFile)
	if err != nil {
		return model.ConfigurationFile{}, err
	}
	for k, c := range configurationFile.Commands {
		for _, ea := range c.ExecArgs {
			values := strings.Split(ea.Value, " ")
			execArgs := make([]model.ExecYamlArg, 0)
			for _, v := range values {
				execArgs = append(execArgs, model.ExecYamlArg{Value: v})
			}
			configurationFile.Commands[k].ExecArgs = execArgs
		}
	}
	return configurationFile, nil
}

func IsValidCommand(c model.YamlCommand) bool {
	if c.Command == "" || c.Exec == "" {
		return false
	}
	return true
}

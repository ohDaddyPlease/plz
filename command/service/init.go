package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ohDaddyPlease/plz/model"
	"github.com/ohDaddyPlease/plz/parser"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Init(extras []interface{}) model.Command {
	c := model.Command{
		Command: "init",
		Use:     `Run 'plz init'' without args`,
		Help:    `Initialize configuration files`,
		Func: func() {
			configContentCommands := make([]model.YamlCommand, 1)
			configContentCommands[0].ExecArgs = make([]model.ExecYamlArg, 1)
			configContentCommands[0].UserArgs = make([]model.UserYamlArg, 1)

			configContent, err := yaml.Marshal(model.ConfigurationFile{
				Commands: configContentCommands,
			})
			if err != nil {
				return
			}
			_, err = os.Stat(parser.ConfigFilePath)
			var needRecreate bool
			if errors.Is(err, os.ErrExist) || err == nil {
				fmt.Println("Configuration file is exists. Do you want to recreate it? [y/n]")
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					if scanner.Text() == "y" {
						needRecreate = true
						break
					} else if scanner.Text() == "n" {
						break
					} else {
						fmt.Println("Please, write 'n' or 'y' for action or CTRL+C/Z for interrupt initialization")
						continue
					}
				}
			} else if errors.Is(err, os.ErrNotExist) {
				needRecreate = true
			}

			if needRecreate {
				config, err := os.Create(parser.ConfigFilePath)
				if err != nil {
					log.Fatal(err)
				}
				defer func() {
					err = config.Close()
					if err != nil {
						log.Fatal(err)
					}
				}()
				_, err = config.Write(configContent)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("File '%s' has been created and stores in '%s'\n", parser.ConfigFileName, parser.ConfigFilePath)
			} else {
				fmt.Println("Configuration file has not been modified")
			}
		},
	}
	return c
}

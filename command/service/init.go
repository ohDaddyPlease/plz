package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ohDaddyPlease/plz/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

func Init() model.Command {
	c := model.Command{
		Name: "init",
		Use: `
Run 'plz init'' without args`,
		Help: `
Initialize configuration files`,
		Func: func() {
			configContent, err := yaml.Marshal(model.ConfigurationFile{
				Commands: make([]model.YamlCommand, 1),
			})
			if err != nil {
				return
			}
			userHomeDir, err := os.UserHomeDir()
			configFileName := ".plz"
			configFilePath := filepath.Join(userHomeDir, configFileName)
			if err != nil {
				log.Fatal(err)
			}
			_, err = os.Stat(configFilePath)
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
				config, err := os.Create(configFilePath)
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

				fmt.Printf("File '%s' has been created and stores in '%s'\n", configFileName, configFilePath)
			} else {
				fmt.Println("Configuration file has not been modified")
			}
		},
	}
	return c
}

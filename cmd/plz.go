package main

import (
	"fmt"
	"github.com/ohDaddyPlease/plz/command/service"
	"github.com/ohDaddyPlease/plz/model"
	"github.com/ohDaddyPlease/plz/parser"
	"log"
	"os"
	"os/exec"
)

var Commands model.CommandType

func init() {
	Commands = make(model.CommandType)

	RegisterWithExtras(service.Version, nil)
	RegisterWithExtras(service.Help, nil)
	RegisterWithExtras(service.Init, nil)
	RegisterWithExtras(service.Commands, []interface{}{Commands})

}

func main() {
	ConfigFileCommands, err := parser.ParseConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range ConfigFileCommands.Commands {
		if !parser.IsValidCommand(c) {
			continue
		}
		var args = make([]model.Arg, len(c.ExecArgs))
		var cmdArgs []string
		for _, a := range c.ExecArgs {
			/*
				if a.Name != "" {
					args = append(args, model.Arg{Name: os.ExpandEnv(a.Name)})
					cmdArgs = append(cmdArgs, os.ExpandEnv(a.Name))
				}
			*/
			if a.Value != "" {
				args = append(args, model.Arg{Name: os.ExpandEnv(a.Value)})
				cmdArgs = append(cmdArgs, os.ExpandEnv(a.Value))
			}

		}

		RegisterWithExtras(func(_ []interface{}) model.Command {
			return model.Command{
				Command: c.Command,
				Use:     c.Use,
				Help:    c.Help,
				Func: func() {
					cmdCommand := exec.Command(c.Exec, cmdArgs...)

					output, err := cmdCommand.Output()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(output))
				},
				Args: args,
			}
		}, nil)
	}

	cmd := parser.ParseArgs()
	c, ok := Commands[cmd.Command]
	if !ok {
		if c.Command.Command != "" {
			fmt.Printf("No command %s .", cmd.Command)
		} else {
			fmt.Print("Please, specify command. ")
		}
		fmt.Println("The list of the possible commands:")
		for _, pc := range Commands {
			fmt.Printf("* plz %s\n", pc.Command.Command)
		}
		os.Exit(0)
	}
	if len(cmd.Args) > 0 {
		switch cmd.Args[0] {
		case "--usage", "-u":
			fmt.Println("[usage]", c.Command.Use)
			os.Exit(0)
		case "--help", "-h":
			fmt.Println("[help]", c.Command.Help)
			os.Exit(0)
		}
	}

	c.Command.Func()
	os.Exit(0)
}

func RegisterWithExtras(cf func(extras []interface{}) model.Command, extras []interface{}) {
	c := cf(extras)
	Commands[c.Command] = model.CommandFields{
		Command: c,
		Extras:  extras,
	}

}

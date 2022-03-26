package main

import (
	"fmt"
	"github.com/ohDaddyPlease/plz/command/service"
	"github.com/ohDaddyPlease/plz/model"
	"github.com/ohDaddyPlease/plz/parser"
	"os"
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

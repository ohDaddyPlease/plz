package main

import (
	"fmt"
	"github.com/ohDaddyPlease/plz/command/service"
	"github.com/ohDaddyPlease/plz/model"
	"github.com/ohDaddyPlease/plz/parser"
	"os"
)

var Commands map[string]model.Command

func init() {
	Commands = make(map[string]model.Command)
	Register(
		service.Version,
		service.Help,
		service.Init,
	)
}

func main() {
	cmd := parser.ParseArgs()
	c, ok := Commands[cmd.Command]
	if !ok {
		fmt.Printf("no command %s\n", cmd.Command)
		os.Exit(0)
	}
	if len(cmd.Args) > 0 {
		switch cmd.Args[0] {
		case "--use", "-u":
			fmt.Println(c.Use)
			os.Exit(0)
		case "--help", "-h":
			fmt.Println(c.Help)
			os.Exit(0)
		}
	}

	c.Func()
	os.Exit(0)
}

func Register(cfs ...func() model.Command) {
	for _, cf := range cfs {
		c := cf()
		Commands[c.Name] = c
	}
}

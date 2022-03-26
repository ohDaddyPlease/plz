package parser

import (
	"github.com/ohDaddyPlease/plz/model"
	"os"
)

func ParseArgs() model.CMD {
	var (
		args    []string
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

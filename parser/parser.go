package parser

import (
	"fmt"
	"github.com/ohDaddyPlease/plz/model"
	"os"
)

var HowToUse = `
	plz [command] ...[args]
`

func ParseArgs() model.CMD {
	if len(os.Args) == 1 {
		fmt.Print(HowToUse)
		os.Exit(1)
	}
	command := os.Args[1]
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	c := model.CMD{
		Command: command,
		Args:    args,
	}
	return c
}

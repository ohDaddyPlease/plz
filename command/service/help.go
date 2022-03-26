package service

import (
	"fmt"
	"github.com/ohDaddyPlease/plz/model"
)

func Help(extras []interface{}) model.Command {
	helpText := `
'plz' is the command-line helper with declarative model in mind`

	c := model.Command{
		Command: "help",
		Use:     helpText,
		Help:    helpText,
		Func: func() {
			fmt.Println(helpText)
		},
	}
	return c
}

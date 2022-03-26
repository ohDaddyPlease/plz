package service

import (
	"fmt"
	"github.com/ohDaddyPlease/plz"
	"github.com/ohDaddyPlease/plz/model"
)

func Version(extras []interface{}) model.Command {
	c := model.Command{
		Command: "version",
		Use: `
Just call 'plz version' to display version without args`,
		Help: `
'plz version' displays current version of the command-line tool`,
		Func: func() {
			fmt.Println("plz command-line tool version is", plz.Version)
		},
	}
	return c
}

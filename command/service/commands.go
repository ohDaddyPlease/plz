package service

import (
	"fmt"
	"github.com/ohDaddyPlease/plz/model"
)

func Commands(extras []interface{}) model.Command {
	c := model.Command{
		Command: "commands",
		Use: `
plz commands`,
		Help: `
Display all possible commands`,
		Func: func() {
			fmt.Println("Possible commands:")
			extra := extras[0]
			e := extra.(model.CommandType)
			for _, c := range e {
				fmt.Printf("* plz %s\n", c.Command.Command)
			}

		},
	}
	return c
}

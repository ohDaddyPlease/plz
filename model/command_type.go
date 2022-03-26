package model

type CommandType map[string]CommandFields
type CommandFields struct {
	Command Command
	Extras  []interface{}
}

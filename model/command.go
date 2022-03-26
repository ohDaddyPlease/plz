package model

type Command struct {
	Name string
	Args []Arg
	Use  string
	Help string
	Func func()
}

type YamlCommand struct {
	Name    string `yaml:"name"`
	Use     string `yaml:"use"`
	Help    string `yaml:"help"`
	Command string `yaml:"command"`
}

type Arg struct {
	Required bool
	Name     string
	Func     func()
}

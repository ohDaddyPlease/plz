package model

type Command struct {
	Command string
	Args    []Arg
	Use     string
	Help    string
	Func    func()
}

type YamlCommand struct {
	Command string    `yaml:"command"`
	Args    []YamlArg `yaml:"args"`
	Use     string    `yaml:"use"`
	Help    string    `yaml:"help"`
	Exec    string    `yaml:"exec"`
}

type Arg struct {
	Required bool
	Name     string
	Func     func()
}

type YamlArg struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

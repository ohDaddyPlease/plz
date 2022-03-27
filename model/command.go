package model

type Command struct {
	Command string
	Args    []Arg
	Use     string
	Help    string
	Func    func()
}

type YamlCommand struct {
	Command  string        `yaml:"command"`
	Exec     string        `yaml:"exec"`
	ExecArgs []ExecYamlArg `yaml:"exec_args"`
	UserArgs []UserYamlArg `yaml:"user_args"`
	Use      string        `yaml:"use"`
	Help     string        `yaml:"help"`
}

type Arg struct {
	Required bool
	Name     string
	Func     func()
}

type ExecYamlArg struct {
	//Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type UserYamlArg struct {
	//Name     string `yaml:"name"`
	Value    string `yaml:"value"`
	Required bool   `yaml:"required"`
}

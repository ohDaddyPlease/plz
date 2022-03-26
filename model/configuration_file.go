package model

type ConfigurationFile struct {
	Commands []YamlCommand `yaml:"commands"`
}

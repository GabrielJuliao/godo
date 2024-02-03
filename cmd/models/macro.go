package models

type Macro struct {
	Name        string `yaml:"name"`
	Executable  string `yaml:"executable"`
	Arguments   string `yaml:"arguments"`
	Description string `yaml:"description"`
}

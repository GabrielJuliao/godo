package main

type Macro struct {
	Name        string `yaml:"name"`
	Executable  string `yaml:"executable"`
	Arguments   string `yaml:"arguments"`
	Description string `yaml:"description"`
}

type Configuration struct {
	Macros []Macro
}

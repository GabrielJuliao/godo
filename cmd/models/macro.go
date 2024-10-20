package models

type Macro struct {
	Name          string `yaml:"name"`
	Executable    string `yaml:"executable"`
	Arguments     string `yaml:"arguments"`
	IsRiskyAction bool   `yaml:"isRiskyAction"`
	Description   string `yaml:"description"`
}

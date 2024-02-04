package configuration

import (
	. "github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/info"
	"github.com/gabrieljuliao/godo/cmd/models"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"regexp"
	"strings"
)

type Configuration struct {
	Macros []models.Macro
}

var ApplicationConfiguration Configuration

func NewConfiguration() *Configuration {

	err := utils.ReadYamlFile(env.Properties[GodoConfigurationFilePath], &ApplicationConfiguration)
	if err != nil {
		if strings.Contains(err.Error(), "yaml: line") {
			log.Println("The configuration file could not be loaded. Please make sure the file is a valid YAML")
			info.PrintValidConfigExample()
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	validateConfiguration(ApplicationConfiguration)

	return &ApplicationConfiguration
}

func validateConfiguration(configuration Configuration) {

	errorCounter := 0

	for _, macro := range configuration.Macros {

		if utils.IsStringEmptyOrNil(macro.Name) {
			log.Println("Macro name cannot be empty")
			errorCounter++
		}

		if !isMacroNameCompliant(macro.Name) {
			log.Println("Macro name must match the following pattern: my-macro-name")
			errorCounter++
		}

		if utils.IsStringEmptyOrNil(macro.Executable) {
			log.Println("Executable cannot be empty")
			errorCounter++
		}

		if utils.IsStringEmptyOrNil(macro.Description) {
			log.Println("Description cannot be empty")
			errorCounter++
		}

	}

	if errorCounter > 0 {
		log.Fatalf("[%d] error(s) were found in your configuration", errorCounter)
	}
}

func isMacroNameCompliant(str string) bool {
	pattern := `^([a-z]+(-[a-z]+)*)?$`
	match, _ := regexp.MatchString(pattern, str)
	return match
}

func ListConfiguration() {
	utils.PrettyPrintYaml(ApplicationConfiguration)
}

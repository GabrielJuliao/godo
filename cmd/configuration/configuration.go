package configuration

import (
	"github.com/gabrieljuliao/godo/cmd/info"
	"github.com/gabrieljuliao/godo/cmd/models"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"os"
	"regexp"
	"strings"
)

type Configuration struct {
	Macros []models.Macro
}

func NewConfiguration() Configuration {
	var config Configuration

	err := utils.ReadYamlFile(findConfigFilePath(), &config)
	if err != nil {
		if strings.Contains(err.Error(), "yaml: line") {
			log.Println("The configuration file could not be loaded. Please make sure the file is a valid YAML")
			info.PrintValidConfigExample()
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	validateConfiguration(config)

	return config
}

// TODO make this madness simple!
func findConfigFilePath() string {
	var filePath = ""

	filePathEnvVar := os.Getenv("GODO_CONFIGURATION_FILE")
	pwd := utils.GetExecutablePath() + string(os.PathSeparator)
	filePathYaml := pwd + "config.yaml"
	filePathYml := pwd + "config.yml"

	if !utils.VerifyFilePath(filePathEnvVar) {
		log.Println("Environment variable GODO_CONFIGURATION_FILE is not set. Falling back to default location(s)")
		switch {
		case utils.VerifyFilePath(filePathYaml):
			log.Printf("Configuration file located at %s", filePathYaml)
			filePath = filePathYaml
		case utils.VerifyFilePath(filePathYml):
			log.Printf("Configuration file located at %s", filePathYml)
			filePath = filePathYml
		default:
			log.Fatalf("Could locate default configuration file(s) [%s, %s]", filePathYaml, filePathYml)
		}
	} else if utils.VerifyFilePath(filePathEnvVar) {
		log.Printf("Configuration file located at %s", filePathEnvVar)
		filePath = filePathEnvVar
	} else {
		log.Fatalf("Could not locate configuration defined in GODO_CONFIGURATION_FILE=%s environment variable.", filePathEnvVar)
	}

	return filePath
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

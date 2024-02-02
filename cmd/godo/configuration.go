package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"regexp"
)

func readConfigurationFile() Configuration {
	var config Configuration
	var filePath string
	filePathEnvVar := os.Getenv("GODO_CONFIGURATION_FILE")
	pwd := getExecutablePath() + string(os.PathSeparator)
	filePathYaml := pwd + "config.yaml"
	filePathYml := pwd + "config.yml"

	if !verifyFilePath(filePathEnvVar) {
		log.Println("Environment variable GODO_CONFIGURATION_FILE is not set. Falling back to default location(s)")
		switch {
		case verifyFilePath(filePathYaml):
			log.Printf("Configuration file located at %s", filePathYaml)
			filePath = filePathYaml
		case verifyFilePath(filePathYml):
			log.Printf("Configuration file located at %s", filePathYml)
			filePath = filePathYml
		default:
			log.Fatalf("Could locate default configuration file(s) [%s, %s]", filePathYaml, filePathYml)
		}
	}

	if verifyFilePath(filePathEnvVar) {
		log.Printf("Configuration file located at %s", filePathEnvVar)
		filePath = filePathEnvVar
	}

	dataBytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Println("The configuration file could not be loaded.")
		log.Fatal(err)
	}

	err = yaml.Unmarshal(dataBytes, &config)

	if err != nil {
		log.Println("Something went wrong when parsing the configuration file, please check if the file is a valid YAML.")
		bStr := "bWFjcm9zOgotIG5hbWU6IGhlbGxvLXdvcmxkCiAgZXhlY3V0YWJsZTogZWNobwogIGFyZ3VtZW50czogIkhlbGxvIFdvcmxkIgogIGRlc2NyaXB0aW9uOiBCcm9hZGNhc3RzICdIZWxsbyBXb3JsZCcgdG8gdGhlIGludGVyZ2FsYWN0aWMgc2hlbGwsIG1ha2luZyBhbGllbnMgb24gRWFydGggYW5kIGluIHNwYWNlIGNyYWNrIGEgY29zbWljIHNtaWxlISA6KQ=="
		log.Printf("Valid configuration file format: \n\n%s\n\n", b64Decoder(bStr))
		log.Fatal(err)
	}

	validateConfiguration(config)

	return config
}

func validateConfiguration(configuration Configuration) {

	errorCounter := 0

	for _, macro := range configuration.Macros {

		if IsStringEmptyOrNil(macro.Name) {
			log.Println("Macro name cannot be empty")
			errorCounter++
		}

		if !isMacroNameCompliant(macro.Name) {
			log.Println("Macro name must match the following pattern: my-macro-name")
			errorCounter++
		}

		if IsStringEmptyOrNil(macro.Executable) {
			log.Println("Executable cannot be empty")
			errorCounter++
		}

		if IsStringEmptyOrNil(macro.Description) {
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

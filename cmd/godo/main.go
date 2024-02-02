package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Macro struct {
	MacroName   string `yaml:"macroName"`
	Executable  string `yaml:"executable"`
	Arguments   string `yaml:"arguments"`
	Description string `yaml:"description"`
}

type Configuration struct {
	Macros []Macro
}

func verifyFilePath(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

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
		bStr := "bWFjcm9zOgotIG1hY3JvTmFtZTogaGVsbG8td29ybGQKICBleGVjdXRhYmxlOiBlY2hvCiAgYXJndW1lbnRzOiAiSGVsbG8gV29ybGQiCiAgZGVzY3JpcHRpb246IEJyb2FkY2FzdHMgJ0hlbGxvIFdvcmxkJyB0byB0aGUgaW50ZXJnYWxhY3RpYyBzaGVsbCwgbWFraW5nIGFsaWVucyBvbiBFYXJ0aCBhbmQgaW4gc3BhY2UgY3JhY2sgYSBjb3NtaWMgc21pbGUhIDop"
		log.Printf("Valid configuration file format: \n\n%s\n\n", b64Decoder(bStr))
		log.Fatal(err)
	}

	return config
}

func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func execMacro(argv []string) {
	config := readConfigurationFile()
	macros := config.Macros

	for _, macro := range macros {

		if macro.MacroName == argv[0] {
			args := strings.Fields(macro.Arguments)

			for _, arg := range argv[1:] {
				args = append(args, arg)
			}
			fmt.Println("")
			fmt.Printf("Macro name: %s\n", macro.MacroName)
			fmt.Printf("Description: %s\n", macro.Description)
			fmt.Printf("Run: %s %s\n\n", macro.Executable, strings.Join(args, " "))
			fmt.Println("")
			execCmd(macro.Executable, args)
			os.Exit(0)
		}

	}
	log.Printf("%s is not a known macro.\n", argv[0])

}

func execCmd(executableName string, arguments []string) {
	cmd := exec.Command(executableName, arguments...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(stdout)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Print(line)
		line, err = reader.ReadString('\n')
	}
}

func b64Decoder(b64Str string) string {
	data, err := b64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func isArgsValid(args []string) bool {
	for _, arg := range args {
		if arg == "godo" {
			log.Fatal("Cannot call godo inside it self.")
		}
	}
	return true
}

func printUsage() {
	fmt.Println("Usage: godo [ macro name ] [ extras arguments ]")
	fmt.Println("Note: [ extras arguments ] will be appended to the end of the string arguments, defined in the configuration file.")
}

func main() {
	fmt.Printf(
		"\n%s\n\n",
		b64Decoder("4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSM4pSA4pSQDQrilIIg4pSs4pSCIOKUgiDilILilILilIIg4pSCDQrilJTilIDilJjilJTilIDilJjilIDilLTilJjilJTilIDilJggdjAuMC4xDQpodHRwczovL2dpdGh1Yi5jb20vR2FicmllbEp1bGlhby9nb2Rv"),
	)
	appArguments := os.Args[1:]
	if len(appArguments) > 0 && isArgsValid(appArguments) && appArguments[0] != "-h" && appArguments[0] != "--help" {
		execMacro(appArguments)
	} else {
		printUsage()
	}
}

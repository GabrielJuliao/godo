package main

import (
	"bufio"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Macro struct {
	MacroName   string
	Executable  string
	Arguments   string
	Description string
}

type Configuration struct {
	Macros []Macro
}

func readConfigurationFile() Configuration {
	var config Configuration
	filePath := getExecutablePath() + string(os.PathSeparator) + "config.json"
	dataBytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Println("The configuration file could not be loaded.")
		log.Fatal(err)
	}

	err = json.Unmarshal(dataBytes, &config)

	if err != nil {
		log.Println("Something went wrong when parsing the configuration file, please check if the file is a valid JSON.")
		bStr := "ew0KICAibWFjcm9zIjogWw0KICAgIHsNCiAgICAgICJtYWNyb05hbWUiOiAiaGVsbG8tY21kIiwNCiAgICAgICJleGVjdXRhYmxlIjogImNtZC5leGUiLA0KICAgICAgImFyZ3VtZW50cyI6ICIvYyBlY2hvIEhlbGxvIFdvcmxkIiwNCiAgICAgICJkZXNjcmlwdGlvbiI6ICJFY2hvIEhlbGxvIFdvcmxkIG1lc3NhZ2UgZnJvbSBjb21tYW5kIHByb21wdCINCiAgICB9DQogIF0NCn0NCg=="
		log.Printf("Configuration file valid format: \n%s\n", b64Decoder(bStr))
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
			fmt.Printf("Macro name: %s\n", macro.MacroName)
			fmt.Printf("Description: %s\n", macro.Description)
			fmt.Printf("Run: %s %s\n\n", macro.Executable, strings.Join(args, " "))
			execCmd(macro.Executable, args)
			os.Exit(0)
		}

	}
	fmt.Printf("%s is not a known macro.\n", argv[0])

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
			log.Fatal("cannot call godo inside it self.")
		}
	}
	return true
}

func main() {
	fmt.Printf(
		"\n%s\n\n",
		b64Decoder("4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSM4pSA4pSQDQrilIIg4pSs4pSCIOKUgiDilILilILilIIg4pSCDQrilJTilIDilJjilJTilIDilJjilIDilLTilJjilJTilIDilJggdjAuMC4xDQpodHRwczovL2dpdGh1Yi5jb20vR2FicmllbEp1bGlhby9nb2Rv"),
	)
	appArguments := os.Args[1:]
	if len(appArguments) > 0 && isArgsValid(appArguments) {
		execMacro(appArguments)
	} else {
		fmt.Println("Usage: godo [ macro name ] [ extras arguments ]")
		fmt.Println("Note: [ extras arguments ] will be appended to the end of the string arguments, defined in the configuration file.")
	}
}

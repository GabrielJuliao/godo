package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func IsStringEmptyOrNil(str any) bool {
	return str == "" || str == nil
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func VerifyFilePath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PrettyPrintYaml(obj any) {
	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(yamlData))
}

func ReadYamlFile(path string, obj any) error {
	dataBytes, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(dataBytes, obj)

	if err != nil {
		return err
	}

	return nil
}

func IsBinaryOnPath(binaryName string) bool {
	_, err := exec.LookPath(binaryName)
	return err == nil
}

func HasNext(list []string, currentIndex int) bool {
	return currentIndex < len(list)-1
}

func PromptForConfirmation(message string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display the prompt message to the user
		fmt.Print(message + " (type 'yes' to confirm, 'no' to cancel): ")

		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Print a blank line for better readability in output
		fmt.Println("")

		// Normalize the input by trimming whitespace and converting to lowercase
		input = strings.TrimSpace(strings.ToLower(input))

		// Check if input is "yes" or "no"
		if input == "yes" {
			return true
		} else if input == "no" {
			return false
		} else {
			// Prompt the user again for invalid input
			fmt.Printf("Invalid input. Please type 'yes' or 'no'.\n\n")
		}
	}
}

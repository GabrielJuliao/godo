package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
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

func IsArgsValid(args []string) bool {
	for _, arg := range args {
		if arg == "godo" {
			log.Fatal("Cannot call godo inside it self.")
		}
	}
	return true
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

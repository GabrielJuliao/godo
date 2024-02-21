package env

import (
	. "github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"os"
	"path"
	"strings"
)

var Properties map[string]string

func InitEnv() *map[string]string {
	Properties = map[string]string{
		GodoConfigurationFilePath:   "./",
		GodoConfigurationEditor:     "",
		GodoConfigurationEditorArgs: "",
	}

	// Override default with OS env
	for key := range Properties {
		if !utils.IsStringEmptyOrNil(os.Getenv(key)) {
			Properties[key] = os.Getenv(key)
		}
	}
	return &Properties
}

func FindConfigFilePath() {

	filePath := Properties[GodoConfigurationFilePath]

	if !strings.HasSuffix(filePath, ".yaml") && !strings.HasSuffix(filePath, ".yml") {
		filePath = strings.TrimSuffix(filePath, string(os.PathSeparator))
		filePath = path.Join(filePath, GodoConfigurationFileName)
	}

	if filePath == GodoConfigurationFileName {
		filePath = utils.GetExecutablePath() + string(os.PathSeparator) + filePath
		log.Printf("%s environment variable is not set. Falling back to default location(s).\n", GodoConfigurationFilePath)
	}

	if utils.VerifyFilePath(filePath) {
		Properties[GodoConfigurationFilePath] = filePath
	} else {
		log.Fatalf("could not locate configuration file at %s", filePath)
	}
}

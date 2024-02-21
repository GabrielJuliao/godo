package exec

import (
	. "github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func OpenConfigurationEditor() {
	var cmd *exec.Cmd

	confEditArgs := strings.Split(env.Properties[GodoConfigurationEditorArgs], ",")
	filePath := env.Properties[GodoConfigurationFilePath]

	confEditArgs = append(confEditArgs, filePath)

	configEditor := env.Properties[GodoConfigurationEditor]

	if utils.IsStringEmptyOrNil(configEditor) {
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", filePath)
		case "linux":
			cmd = exec.Command("xdg-open", filePath)
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		default:
			log.Println("The default configuration does not support this OS.")
			log.Fatal(os.ErrNotExist)
		}
	} else {
		log.Printf("Edit command: %s %s", configEditor, confEditArgs)
		cmd = exec.Command(configEditor, confEditArgs...)
	}

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

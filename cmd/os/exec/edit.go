package exec

import (
	. "github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func OpenConfigurationEditor() {
	var cmd *exec.Cmd
	filePath := env.Properties[GodoConfigurationFilePath]

	if utils.IsStringEmptyOrNil(env.Properties[GodoConfigurationEditor]) {
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", filePath)
		case "linux":
			cmd = exec.Command("xdg-open", filePath)
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		default:
			log.Fatal(os.ErrNotExist)
		}
	} else {
		cmd = exec.Command(env.Properties[GodoConfigurationEditor], env.Properties[GodoConfigurationEditorArgs], filePath)
	}

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

package exec

import (
	. "github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"os"
	"runtime"
	"strings"
)

func OpenConfigurationEditor() {
	var args []string
	filePath := env.Properties[GodoConfigurationFilePath]
	configEditor := env.Properties[GodoConfigurationEditor]

	if utils.IsStringEmptyOrNil(configEditor) {
		switch runtime.GOOS {
		case "darwin":
			args = append(args, filePath)
			Cmd("open", args)
		case "linux":
			args = append(args, filePath)
			editor := findLinuxDefaultTextEditor()
			if !utils.IsStringEmptyOrNil(editor) {
				Cmd(editor, args)
			} else {
				log.Fatalf("Could not determine the default text editor. Please provide the default text editor either by using the option flags or the environment variables.")
			}
		case "windows":
			args = []string{"/c", "start", filePath}
			Cmd("cmd", args)
		default:
			log.Println("Unsupported Operating System.")
			log.Fatal(os.ErrNotExist)
		}
	} else {
		customArgs := env.Properties[GodoConfigurationEditorArgs]
		if !utils.IsStringEmptyOrNil(customArgs) {
			args = append(args, strings.Split(customArgs, ",")...)
		}
		args = append(args, filePath)
		Cmd(configEditor, args)
	}
}

func findLinuxDefaultTextEditor() string {
	textEditors := []string{"xdg-open", "nano", "vi"}

	for _, editor := range textEditors {
		if utils.IsBinaryOnPath(editor) {
			return editor
		}
	}
	return ""
}

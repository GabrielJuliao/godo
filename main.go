package main

import (
	_ "embed"
	"github.com/gabrieljuliao/godo/cmd/configuration"
	"github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/info"
	"github.com/gabrieljuliao/godo/cmd/os/exec"
	"github.com/gabrieljuliao/godo/cmd/service"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//go:embed VERSION
var appVersion string

//go:embed resources/banner.txt
var appBanner string

//go:embed resources/usage.txt
var appUsageMsg string

//go:embed resources/config_file_example.yaml
var appConfigFileExample string

func main() {
	info.NewAppInfo(appVersion, appUsageMsg, appBanner, appConfigFileExample)
	info.PrintBanner()
	env.InitEnv()

	preConfigActions(filterGodoArgs(os.Args[1:]))
}

func filterGodoArgs(args []string) []string {
	for i, arg := range args {
		if strings.HasPrefix(arg, "--godo") {
			if !utils.HasNext(args, i) {
				log.Fatal("You must provide a value to godo options.")
			}
			noPrefixArg := strings.Replace(arg, "--godo-", "", -1)
			switch noPrefixArg {
			case "config-file":
				path := args[i+1]
				if filepath.IsAbs(path) {
					env.Properties[consts.GodoConfigurationFilePath] = path
					args = args[2:]
				} else {
					log.Fatalf("Value '%s' for option '%s' is not a valid path.", path, arg)
				}
			case "config-editor":
				path := args[i+1]
				if utils.IsBinaryOnPath(path) {
					env.Properties[consts.GodoConfigurationEditor] = args[i+1]
					args = args[2:]
				} else {
					log.Fatalf("Could not locate '%s' for option '%s'. Make sure executable is exported to path.", path, arg)
				}
			case "config-editor-args":
				regex, _ := regexp.Compile("^(?:[^,\\s]+(?:,[^,\\s]+)*)?$\n")
				editorArgs := args[i+1]
				if regex.MatchString(editorArgs) {
					env.Properties[consts.GodoConfigurationEditorArgs] = args[i+1]
					args = args[2:]
				} else {
					log.Fatalf("Value '%s' for option '%s' does not match the pattern: --arg1,value1,arg-2.", editorArgs, arg)
				}
			default:
				log.Fatalf("Option '%s' does not exist", arg)
			}
		}
	}
	return args
}

func preConfigActions(argv []string) {
	env.FindConfigFilePath()
	var action = ""

	if len(argv) > 0 {
		action = argv[0]
	}

	switch {
	case action == "edit":
		exec.OpenConfigurationEditor()
	case action == "" || action == "-h" || action == "--help":
		info.PrintUsage()
	default:
		postConfigActions(action, argv)
	}
}

func postConfigActions(action string, argv []string) {
	// load configuration
	configuration.NewConfiguration()

	switch action {
	case "list":
		configuration.ListConfiguration()
	default:
		service.ExecMacro(argv)
	}

}

package main

import (
	_ "embed"
	"github.com/gabrieljuliao/godo/cmd/configuration"
	"github.com/gabrieljuliao/godo/cmd/consts"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/info"
	"github.com/gabrieljuliao/godo/cmd/os/exec"
	"github.com/gabrieljuliao/godo/cmd/service"
	"os"
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
	env.FindConfigFilePath()

	preConfigActions(filterGodoArgs(os.Args[1:]))
}

func filterGodoArgs(args []string) []string {
	for i, arg := range args {
		if strings.HasPrefix(arg, "--godo") {
			arg = strings.Replace(arg, "--godo-", "", -1)
			switch arg {
			case "config-file":
				env.Properties[consts.GodoConfigurationFilePath] = args[i+1]
				args = args[2:]
			case "config-editor":
				env.Properties[consts.GodoConfigurationEditor] = args[i+1]
				args = args[2:]
			case "config-editor-args":
				env.Properties[consts.GodoConfigurationEditorArgs] = args[i+1]
				args = args[2:]
			}
		}
	}
	return args
}

func preConfigActions(argv []string) {
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

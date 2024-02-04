package main

import (
	_ "embed"
	"github.com/gabrieljuliao/godo/cmd/configuration"
	"github.com/gabrieljuliao/godo/cmd/env"
	"github.com/gabrieljuliao/godo/cmd/info"
	"github.com/gabrieljuliao/godo/cmd/os/exec"
	"github.com/gabrieljuliao/godo/cmd/service"
	"os"
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
	env.InitEnv()
	env.FindConfigFilePath()
	info.NewAppInfo(appVersion, appUsageMsg, appBanner, appConfigFileExample)
	info.PrintBanner()
	preConfigActions()
}

func preConfigActions() {
	var action = ""
	argv := os.Args[1:]

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

package main

import (
	_ "embed"
	"github.com/gabrieljuliao/godo/cmd/context"
	"github.com/gabrieljuliao/godo/cmd/info"
	"github.com/gabrieljuliao/godo/cmd/service"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"os"
)

func main() {
	start()
	choice()
}

func choice() {
	var action = ""
	argv := os.Args[1:]

	if len(argv) > 0 {
		action = argv[0]
	}

	switch action {
	case "list":
		context.ListConfiguration()
	default:
		executeMacro(argv)
	}

}

func executeMacro(appArguments []string) {
	if len(appArguments) > 0 && utils.IsArgsValid(appArguments) && appArguments[0] != "-h" && appArguments[0] != "--help" {
		service.ExecMacro(appArguments)
	} else {
		info.PrintUsage()
	}
}

//go:embed VERSION
var appVersion string

//go:embed resources/banner.txt
var appBanner string

//go:embed resources/usage.txt
var appUsageMsg string

//go:embed resources/config_file_example.yaml
var appConfigFileExample string

func start() {
	info.NewAppInfo(appVersion, appUsageMsg, appBanner, appConfigFileExample)
	info.PrintBanner()
	context.NewContext()
}

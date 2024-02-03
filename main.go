package main

import (
	_ "embed"
	"github.com/gabrieljuliao/godo/cmd/context"
	"github.com/gabrieljuliao/godo/cmd/service"
	"github.com/gabrieljuliao/godo/cmd/utils"
	"log"
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
	context.NewAppInfo(appVersion, appUsageMsg, appBanner, appConfigFileExample)
	utils.PrintBanner()
	appArguments := os.Args[1:]
	if len(appArguments) > 0 && isArgsValid(appArguments) && appArguments[0] != "-h" && appArguments[0] != "--help" {
		service.ExecMacro(appArguments)
	} else {
		utils.PrintUsage()
	}
}

func isArgsValid(args []string) bool {
	for _, arg := range args {
		if arg == "godo" {
			log.Fatal("Cannot call godo inside it self.")
		}
	}
	return true
}

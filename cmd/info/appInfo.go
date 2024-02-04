package info

import (
	"fmt"
	"log"
	"strings"
)

type ApplicationInformation struct {
	Version           string
	UsageMessage      string
	Banner            string
	ConfigFileExample string
}

var AppInfo ApplicationInformation

func NewAppInfo(version string, usageMessage string, banner string, configFileExample string) *ApplicationInformation {
	AppInfo = ApplicationInformation{
		version,
		usageMessage,
		buildBanner(banner, version),
		configFileExample,
	}
	return &AppInfo
}

func PrintUsage() {
	fmt.Printf("\n%s\n", AppInfo.UsageMessage)
}

func PrintBanner() {
	fmt.Printf("\n%s\n\n", AppInfo.Banner)
}

func PrintValidConfigExample() {
	log.Printf("Valid configuration file format: \n\n%s\n\n", AppInfo.ConfigFileExample)
}

func buildBanner(banner string, version string) string {
	return strings.Replace(banner, "<VERSION>", version, -1)
}

package context

import "strings"

type AppInfo struct {
	Version           string
	UsageMessage      string
	Banner            string
	ConfigFileExample string
}

var ApplicationInfo AppInfo

func NewAppInfo(version string, usageMessage string, banner string, configFileExample string) *AppInfo {
	ApplicationInfo = AppInfo{
		version,
		usageMessage,
		buildBanner(banner, version),
		configFileExample,
	}
	return &ApplicationInfo
}

func buildBanner(banner string, version string) string {
	return strings.Replace(banner, "<VERSION>", version, -1)
}

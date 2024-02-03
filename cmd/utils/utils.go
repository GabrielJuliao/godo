package utils

import (
	"fmt"
	"github.com/gabrieljuliao/godo/cmd/context"
	"log"
	"os"
	"path/filepath"
)

func IsStringEmptyOrNil(str any) bool {
	return str == "" || str == nil
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func VerifyFilePath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PrintUsage() {
	fmt.Println(context.ApplicationInfo.UsageMessage)
}

func PrintBanner() {
	fmt.Printf("\n%s\n\n", context.ApplicationInfo.Banner)
}

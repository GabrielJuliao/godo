package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func IsStringEmptyOrNil(str any) bool {
	return str == "" || str == nil
}

func b64Decoder(b64Str string) string {
	data, err := b64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func verifyFilePath(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

func printUsage() {
	fmt.Println("Usage: godo [ macro name ] [ extras arguments ]")
	fmt.Println("Note: [ extras arguments ] will be appended to the end of the string arguments, defined in the configuration file.")
}

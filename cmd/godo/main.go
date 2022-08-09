package main

import (
	"bufio"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Macro struct {
	MacroName   string
	Executable  string
	Arguments   string
	Description string
}

type Configuration struct {
	Macros []Macro
}

func readConfigurationFile() Configuration {
	var config Configuration
	filePath := getExecutablePath() + string(os.PathSeparator) + "config.json"
	dataBytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(dataBytes, &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func execMacro(argv []string) {
	config := readConfigurationFile()
	macros := config.Macros

	for _, macro := range macros {

		if macro.MacroName == argv[0] {
			args := strings.Fields(macro.Arguments)

			for _, arg := range argv[1:] {
				args = append(args, arg)
			}

			execCmd(macro.Executable, args)
			os.Exit(0)
		}

	}
	fmt.Printf("\n%s is not a known macro\n", argv)

}

func execCmd(executableName string, arguments []string) {
	cmd := exec.Command(executableName, arguments...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(stdout)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Print(line)
		line, err = reader.ReadString('\n')
	}
}

func printInit() {
	b64str := "ICAgX19fXyAgICBfX18gICAgX19fXyAgICAgX19fICANCiAgLyBfX198ICAvIF8gXCAgfCAgXyBcICAgLyBfIFwgDQogfCB8ICBfICB8IHwgfCB8IHwgfCB8IHwgfCB8IHwgfA0KIHwgfF98IHwgfCB8X3wgfCB8IHxffCB8IHwgfF98IHwNCiAgXF9fX198ICBcX19fLyAgfF9fX18vICAgXF9fXy8gIHYwLjAuMQ0KDQpodHRwczovL2dpdGh1Yi5jb20vR2FicmllbEp1bGlhby9nb2Rv"
	data, err := b64.StdEncoding.DecodeString(b64str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(data))
}

func main() {
	printInit()
	argv := os.Args[1:]
	argc := len(argv)

	if argc > 0 {
		execMacro(argv)
	} else {
		fmt.Printf("\nUsage: godo [macro name] [arguments]\n")
	}
}

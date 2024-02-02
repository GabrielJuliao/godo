package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func execMacro(argv []string) {
	config := readConfigurationFile()
	macros := config.Macros

	for _, macro := range macros {

		if macro.Name == argv[0] {
			args := strings.Fields(macro.Arguments)

			for _, arg := range argv[1:] {
				args = append(args, arg)
			}
			fmt.Println("")
			fmt.Printf("Macro name: %s\n", macro.Name)
			fmt.Printf("Description: %s\n", macro.Description)
			fmt.Printf("Run: %s %s\n\n", macro.Executable, strings.Join(args, " "))
			fmt.Println("")
			execCmd(macro.Executable, args)
			os.Exit(0)
		}

	}
	log.Printf("%s is not a known macro.\n", argv[0])

}

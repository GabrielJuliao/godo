package service

import (
	"fmt"
	"github.com/gabrieljuliao/godo/cmd/configuration"
	"github.com/gabrieljuliao/godo/cmd/os/exec"
	"log"
	"os"
	"strings"
)

func ExecMacro(argv []string) {
	verifyArgs(argv)
	for _, macro := range configuration.ApplicationConfiguration.Macros {

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
			exec.Cmd(macro.Executable, args)
			os.Exit(0)
		}

	}
	log.Printf("%s is not a known macro.\n", argv[0])
}

func verifyArgs(args []string) {
	for _, arg := range args {
		if arg == "godo" {
			log.Fatal("Cannot call godo inside it self.")
		}
	}
}

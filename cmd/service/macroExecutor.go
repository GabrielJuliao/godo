package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/gabrieljuliao/godo/cmd/configuration"
	"github.com/gabrieljuliao/godo/cmd/os/exec"
	"github.com/gabrieljuliao/godo/cmd/utils"
)

func ExecMacro(argv []string) {
	verifyArgs(argv)
	for _, macro := range configuration.ApplicationConfiguration.Macros {

		if macro.Name == argv[0] {
			args := strings.Fields(macro.Arguments)

			args = append(args, argv[1:]...)

			fmt.Println("")
			fmt.Printf("Macro name: %s\n", macro.Name)
			fmt.Printf("Description: %s\n", macro.Description)
			fmt.Printf("Command(s): %s %s\n\n", macro.Executable, strings.Join(args, " "))

			if macro.IsRiskyAction {
				if !utils.PromptForConfirmation("WARNING: This Macro is potentially dangerous and has been set to require confirmation before execution.\nAre you sure you want to continue?") {
					return
				}
			}

			exec.Cmd(macro.Executable, args)
			return
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

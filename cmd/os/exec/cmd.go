package exec

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Cmd(executableName string, arguments []string) {
	cmdStrMsg := executableName + " " + strings.Join(arguments, " ")
	log.Printf("Executing command: %s", cmdStrMsg)

	cmd := exec.Command(executableName, arguments...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Error while executing: %s", cmdStrMsg)
		log.Fatal(err)
	}
}

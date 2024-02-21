package exec

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Cmd(executableName string, arguments []string) {
	cmd := exec.Command(executableName, arguments...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			_, err = fmt.Fprint(os.Stderr, line)
			if err != nil {
				return
			}
		}
	}()

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println("\nGODO: Error waiting for command to finish:", err)
	}
}

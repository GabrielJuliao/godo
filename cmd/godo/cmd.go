package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

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

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command to finish:", err)
	}
}

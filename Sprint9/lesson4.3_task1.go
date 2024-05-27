package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("echo", "Hello, world!")
	stdout, err := echoCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	catCmd := exec.Command("cat")
	catCmd.Stdout = os.Stdout
	catCmd.Stdin = stdout
	if err = echoCmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err = catCmd.Start(); err != nil {
		log.Fatal(err)
	}

	echoCmd.Wait()
	catCmd.Wait()
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// cmd := exec.Command("echo", "-n", "{\"Class\": \"verbe\", \"Frequency\":351960, \"Label\": \"être\"}")
	cmd := exec.Command("curl", `"https://la-tim.fr"`)

	// Get the stdout pipe
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Read the output
	var output bytes.Buffer
	_, err = io.Copy(&output, stdout)
	if err != nil {
		fmt.Println("Error reading stdout:", err)
		return
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
		return
	}

	// Print the output
	fmt.Println("Command output:", output.String())
}

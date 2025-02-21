package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func main() {

	// curl := "curl"
	// link := "\"https://la-tim.fr\""
	// data_option := "--data-raw 'word=négatif'"

	cmd := exec.Command(`curl`)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.CmdLine = `curl "https://la-tim.fr"`

	// fmt.Println("'cakah'")
	// tail := "tail"
	// cem := "cem.go"
	// opt := "-n 3"

	// cmd := exec.Command(tail, cem, opt)
	// stdout, err := cmd.Output()

	// cmd := exec.Command(curl, 	)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

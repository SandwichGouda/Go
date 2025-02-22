package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	cmd := exec.Command("echo", "-n", "{\"Class\": \"verbe\", \"Frequency\":351960, \"Label\": \"être\"}")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var word struct {
		Class     string
		Frequency int
		Label     string
	}
	if err := json.NewDecoder(stdout).Decode(&word); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(word.Class, word.Frequency, word.Label)
}

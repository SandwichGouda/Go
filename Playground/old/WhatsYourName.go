package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your name? ")
	name, _ := reader.ReadString('\n')
	// Trim the newline character from the input
	name = name[:len(name)-1]
	fmt.Printf("Nice to meet you, %s!\n", name)
}

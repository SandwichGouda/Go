package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How old are you ? ")
	nb, _ := reader.ReadString('\n')
	nb = nb[:len(nb)-1]

	day, err := strconv.Atoi(nb)

	for err != nil {
		fmt.Print("Input convertion to int type failed. Please type in the puzzle day number : ")
		nb, _ := reader.ReadString('\n')
		nb = nb[:len(nb)-1]
	}
	fmt.Println(day)
}

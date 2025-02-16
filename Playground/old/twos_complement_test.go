package main

import "fmt"

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * pow(x, n-1)
}

func main() {

	k := 60
	for k < 70 {
		fmt.Printf("2**%d : %d \n", k, pow(2, k))
		k++
	}

	/*
		a := 3
		for a < pow(2, 63) {
			a *= 2
			fmt.Printf("a : %d\n", a)
			}
	*/
}

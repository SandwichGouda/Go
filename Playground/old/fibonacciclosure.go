package main

import "fmt"

func add(x int, y int) (z int) {
	z = x + y
	return
}

func f(x int) int {
	return x + 1
}

type TwoDIntVector struct {
	X int
	Y int
}

func Fibonacci() func() int {
	a := 0
	b := 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()
	for k := 0; k <= 15; k++ {
		fmt.Println(f())
	}
}

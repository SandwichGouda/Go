package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func incr(x *int) {
	*x++
}

type myint int

func (x myint) incr(y myint) {
	x = x + y
}

func main() {
	x := myint(2)
	x.incr(myint(2))
	fmt.Println(x)
}

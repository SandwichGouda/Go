package main

import (
	"fmt"
	"math"
)

/* Declare a single variable */
var a int

/* Declare a group of variables */
var (
	b bool
	c float32
	d string
)

func add(x int, y int) int {
	return x + y
}

func mul(x, y int, s string) int {
	fmt.Println("Multiplying", x, "and", y, "together. Additional message :", s)
	return x * y
}

func swap(x int, s string) (string, int) {
	return s, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum_of_squares(n int) int {
	s := 0
	for k := 1; k <= n; k++ {
		s += k * k
	}
	return s
}

func useless_program_1() {
	sum := 0
	for sum < 1000 {
		sum += sum
	}
}

func useless_program_2() {
	sum := 0
	for sum < 1000 {
		sum += sum
	}
}

func trap() {
	for {
	}
}

func iseven(n int) bool {
	if math.Mod(float64(n), 2) == 0 {
		return true
	} else {
		return false
	}
}

func is_1_or_0(n int) bool {
	switch p := n; p {
	case 0:
		fmt.Println("aa")
		return true
	case 1:
		return true
	default:
		return false
	}
}
func main() {
	/* Assigning declared variables */
	a = 42            // Assign single value
	b, c = true, 32.0 // Assign multiple values
	d = "string"      // Strings need double quotes

	/* Declaring and assigning one variable int one go */
	e := 5
	var f = 5
	var g int = 5

	/* Declaring and assigning multiple variables, all in one go */
	h, i := true, "indeed"
	var j, k = false, 5
	var l, m bool = false, true // Here variables must have the same type

	var (
		a int = 5
		b     = "a"
	)

	/* Declaring consts */
	const n = 5
	const o int = 5
	const p, q int = 5, 6

	const (
		pc, qc = 5, 6
	)

	r := float64(p)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, pc, qc, r, add(p, q), mul(p, q, "nothing"))
	fmt.Println(math.Pi)
	fmt.Println(r, "\n")
	for n := 0; n <= 10; n++ {
		fmt.Println(sum_of_squares(n))
	}
	// fmt.Println(math.pi) // Doesn't work
	is_1_or_0(0)
}

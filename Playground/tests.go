package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1.
	for k := 0; x > 0; k++ {
		x = math.Pow(math.E, float64(-k))
		fmt.Println(k, x)
	}
}

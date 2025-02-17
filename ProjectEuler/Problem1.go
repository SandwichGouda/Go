package main

import (
	"fmt"
	"math"
)

func main() {
	s := 0

	for k := 0; 3*k < 1000; k++ {
		s += 3 * k
	}

	for k := 0; 5*k < 1000; k++ {
		if math.Mod(float64(5*k), 3) > 0 {
			s += 5 * k
		}
	}

	fmt.Println(s)
}

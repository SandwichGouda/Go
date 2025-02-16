package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for k := 0; k <= 1000; k++ {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return strings.Join([]string{"cannot Sqrt negative number:}", strconv.Itoa(-2)}, " ")
	} else {
		return ""
	}
}

func Sqrt(x float64) (float64, error) {
	error := x.Error()
}

func main() {
	fmt.Println(Sqrt(2))
}

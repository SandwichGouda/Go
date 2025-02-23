package main

import (
	"fmt"
)

type TwoDVector struct {
	X int
	Y int
}

type ThreeDVector struct {
	X int
	Y int
	Z int
}

func Proj(th ThreeDVector) TwoDVector {
	var tw TwoDVector

	th.X = tw.X
	th.Y = tw.Y

	return tw
}

func main() {
	v := ThreeDVector{X: 3, Y: 2, Z: 1}
	fmt.Println(Proj(v))
}

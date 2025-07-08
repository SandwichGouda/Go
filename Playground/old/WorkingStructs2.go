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

func Proj(th *ThreeDVector) *TwoDVector {

	return &TwoDVector{
		X: th.X,
		Y: th.Y,
	}
}

func main() {
	v := &ThreeDVector{X: 3, Y: 2, Z: 1}
	fmt.Println(Proj(v))
}

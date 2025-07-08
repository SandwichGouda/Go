package main

import (
	"fmt"
	"strconv"
)

func main1() {

	var t = []int{1,2,3}
	var x = make([]int, 10000)

	x = append(x,t...)

	fmt.Println(t,x)
}


func main() {

	var x = make([]string, 10000)

	for k := 0 ; k < 10000 ; k++ {
		x[k] = strconv.Itoa(k)
	}

	fmt.Println(x)
}

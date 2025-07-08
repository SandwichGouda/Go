package main

import (
	"fmt"
	"os"
	"strconv"
)

func pascalsTriangle(n int) [][]int {
	if (n <= 0) {
		return make([][]int, 0)
	}
	pT := make([][]int, n)
	for i := range n {
		pT[i] = make([]int, n)
	}
	pT[0][0] = 1
	for i := 1 ; i < n ; i++ {
		pT[i][0] = 1
		for j := 1 ; j < n ; j++ {
			pT[i][j] = pT[i-1][j] + pT[i-1][j-1]
		}
	}
	return pT
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	pT := pascalsTriangle(n)
}
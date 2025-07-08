package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
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
	// n choose k : pT[n][k]

	kCondition := func(k int, seq []int) bool { 
		// The zero-th condition is well defined (and should be always verified, a_0 = 0)
		l := len(seq)
		if k >= l {
			log.Fatalf("k >= len(seq) in kCondition (k, seq)",k,seq)
		}
		S := 0
		pow := 1
		for j := range(k+1) {
			S += pT[k][seq[j]]
			// fmt.Println(k,seq[j],pT[k][seq[j]])
			pow *= 2
		}
		pow /= 2
		// fmt.Println(pow,S)
		return S == pow
	}
	
	worksUntil := func(seq []int) int {
		// Returns the first index of seq that makes the sequence fail to check the conditions. 
		// Returns len(seq) (last index + 1) if cheks all conditions.
		k := 0
		for (k < len(seq) && kCondition(k, seq)) {
			k++
		}
		return k
	}

	seq := make([]int, n)	

	for worksUntil(seq) < n {
		seq[worksUntil(seq)]++
		
	}
	fmt.Println(seq)
	fmt.Println(worksUntil(seq) == n)
}
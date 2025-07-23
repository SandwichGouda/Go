package main

import (
	"math/rand"
	"fmt"
	"os"
	"strconv"
	"log"
)

func createMap(n int) [][]int {
	matrix := make([][]int, n)

	for i := range(n) {
		matrix[i] = make([]int, n)
	}

	for i := range n {
		for j := range i {
			matrix[i][j] = rand.Intn(2)*2 - 1
		}
	}

	for j := range n {
		for i := range j {
			matrix[i][j] = - matrix[j][i]
		}
	}
	
	return matrix
}

func randomOrder(n int) [][2]int {
	order := make([][2]int,n*(n-1)/2)

	k := 0
	for j := range n {
		for i := range j {
			order[k][0] = i
			order[k][1] = j
			k++
		}
	}

	rand.Shuffle(len(order), func(i, j int) {
		order[i], order[j] = order[j], order[i]
	})

	return order
}

func isDefinitelyNotOneWayOut(k int, partialMap [][]int) bool {
	n := len(partialMap)

	// Check if it has at least 2 outputs
	t := 0
	for i := range n {
		if (partialMap[i][k] == 1) {
			t += 1
		}
	}

	if (t >= 2) {
		return true
	}

	// Check if it has at least n-2 inputs
	t = 0
	for i := range n {
		if (partialMap[i][k] == -1) {
			t += 1
		}
	}

	if (t >= n-2) {
		return true
	}

	return false
}

func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Printf("[")
		for _, val := range row {
			// Print each value with the same width
			fmt.Printf("%*d", 3, val)
		}
		fmt.Printf(" ]")
		fmt.Println()
	}
}

func applyAlgorithm(matrix [][]int, order [][2]int) int {
	n := len(matrix)

	partialMap := make([][]int,n)
	for i := range n {
		partialMap[i] = make([]int, n)
	}

	var i,j int
	q := 0

	// First phase
	for _,v := range(order) {
		i = v[0]
		j = v[1]

		if !isDefinitelyNotOneWayOut(i, partialMap) && !isDefinitelyNotOneWayOut(j, partialMap) {
			partialMap[i][j] = matrix[i][j]
			partialMap[j][i] = matrix[j][i]
			q++
		}
	}

	// fmt.Println("Partial map after first phase:")
	// displayMatrix(partialMap)

	// Second phase
	for _,v := range(order) {
		i = v[0]
		j = v[1]

		if !isDefinitelyNotOneWayOut(i, partialMap) || !isDefinitelyNotOneWayOut(j, partialMap) {
			partialMap[i][j] = matrix[i][j]
			partialMap[j][i] = matrix[j][i]
			q++
		}
	}

	// fmt.Println("Partial map after second phase:")
	// displayMatrix(partialMap)

	return q
}

func main() {
	var n,seed int
	if len(os.Args) == 1 {
		n = 5
	} else if len(os.Args) == 2 {
		n,_ = strconv.Atoi(os.Args[1])
	} else if len(os.Args) == 3 {
		n,_ = strconv.Atoi(os.Args[1])
		seed,_ = strconv.Atoi(os.Args[2])
		rand.Seed(int64(seed))
	} else {
		log.Fatal("Too many arguments. The first argument (optional) should be the desired size n (default 5) and the second (optional and only if first argument is given) the random seed (int).")
	}

	order := randomOrder(n)
	// fmt.Println(order)

	matrix := createMap(n)

	// fmt.Println("Map")
	// displayMatrix(matrix)

	// fmt.Println("q = ",applyAlgorithm(matrix,order))
	// fmt.Println("4n = ",4*n)

	q := applyAlgorithm(matrix,order)
	if 4*n-q <= 3 {
		fmt.Print(4*n-q, " ")
	}
}
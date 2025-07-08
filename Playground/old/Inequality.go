package main

import (
	"math/rand"
	"math"
	"fmt"
)

func main() {

	var (
		n int
		// eps []float64
		logsum float64
		sum float64
		max float64
		e float64
	)
	
	
	for {
		n = rand.Intn(100)+2
		// eps = make([]int, n)

		logsum = 0
		sum = 0
		for k := 0 ; k < n ; k++ {
			e = rand.Float64()-1
			// eps[k] = 
			sum += e
			logsum += -math.Log(1-e)
			if e > max {
				max = e 
			}
		}

		if (logsum <= sum/(1-max)) && (sum < 0) {
			fmt.Println("uh oh")
			break
		} else {
			fmt.Print("yes")
		}
	}
}

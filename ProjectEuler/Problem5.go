package main

import "fmt"

func SmallestEvenlyDivisible(n int) int {

	if n == 1 {
		return 1
	}

	m := n
	b := false

	for !b {
		for k := 2 ; k <= n ; k++ {
			if (m % k > 0) {
				break
			}
			if (k == n) {
				b = true
				m--
				break
			}
		}
		m++
	}

	return m
}


func main() {
	fmt.Println(SmallestEvenlyDivisible(20))
}
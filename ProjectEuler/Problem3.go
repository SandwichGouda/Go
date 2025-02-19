package main

import (
	"fmt"
	"math"
	"slices"
)

// Returns a slice of all primes less than or equal to n using Eratosthene's sieve.
func Eratosthene(n int) []int {

	isprime := slices.Repeat([]bool{true}, n+1)

	isprime[0] = false
	isprime[1] = false

	var j int
	for k := 2; k <= n; k++ {
		j = 2
		for k*j <= n {
			isprime[k*j] = false
			j++
		}

	}

	var primes []int

	for k := 0; k <= n; k++ {
		if isprime[k] {
			primes = append(primes, k)
		}
	}

	return primes
}

func Isprime(n int) bool {

	for k := 2; k < n; k++ {
		if math.Mod(float64(n), float64(k)) == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 600851475143
	m := 0
	for k := 2; k < n; k++ {
		if Isprime(k) {
			if math.Mod(float64(n), float64(k)) == 0 {
				m = k
			}
		}
	}
	fmt.Println(m)
}

/*
func main() {
	n := 600851475143

	primes := Eratosthene(n / 2)

	l := len(primes)

	for p := l - 1; p >= 0; p-- {
		if math.Mod(float64(n), float64(primes[p])) == 0 {
			fmt.Println(primes[p])
			break
		}
	}

}
*/

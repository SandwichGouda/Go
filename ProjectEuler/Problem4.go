package main

import (
	"fmt"
	"math"
)

func Pow(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * Pow(x, n-1)
}

func NbDigits(n int) int {
	if n == 1 {
		return 1
	}
	digits := int(math.Log10(float64(n))) + 1

	// Special case if n is a power of 10
	if n == Pow(10, digits-1) {
		return digits - 1
	}
	return digits
}

func Mirror(n int) int {
	nb := NbDigits(n)
	var d int
	var m int

	for i := 0; i <= nb-1; i++ {
		d = n % 10
		m += Pow(10, nb-1-i) * d
		n = (n - d) / 10
	}

	return m
}

func IsPalindrome(n int) bool {
	return n == Mirror(n)
}

func LargestPalindrome(n int) (m int) {

	for a := Pow(10, n-1); a < Pow(10, n); a++ {
		for b := Pow(10, n-1); b < Pow(10, n); b++ {
			if a*b > m {
				if IsPalindrome(a * b) {
					m = a * b
				}
			}
		}
	}
	return
}

func Reverse(s []int) []int {

	l := len(s)

	r := make([]int, l, l)

	for k := 0; k < l; k++ {
		r[l-k-1] = s[k]
	}
	return r
}

func IntToSlice(n int) []int {
	nb := NbDigits(n)
	var d int
	sl := make([]int, nb, nb)

	for i := 0; i <= nb-1; i++ {
		d = n % 10
		sl[nb-i-1] = d
		n = (n - d) / 10
	}

	return sl
}

func IsPalindrome2(n int) (b bool) {
	b = true
	sl := IntToSlice(n)
	r := Reverse(sl)
	l := len(sl)

	for k := 0; k < l; k++ {
		if sl[k] != r[k] {
			b = false
		}
	}
	return
}

func LargestPalindrome2(n int) (m int) {

	for a := Pow(10, n-1); a < Pow(10, n); a++ {
		for b := Pow(10, n-1); b < Pow(10, n); b++ {
			if a*b > m {
				if IsPalindrome2(a * b) {
					m = a * b
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(LargestPalindrome2(3))
}

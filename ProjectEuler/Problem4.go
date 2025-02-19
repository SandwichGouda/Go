package main

import (
	"fmt"
	"math"
)

func Pow(x int, n uint) int {
	if n == 0 {
		return 1
	}
	return x * Pow(x, n-1)
}

func NbDigits(n uint) uint {
	if n == 1 {
		return 1
	}
	digits := uint(math.Log10(float64(n))) + 1

	// Special case if n is a power of 10
	if int(n) == Pow(10, digits-1) {
		return digits - 1
	}
	return digits
}

func Mirror(n uint) uint {
	nb := NbDigits(n)
	var d, m uint

	for i := uint(0); i <= nb-1; i++ {
		d = n % 10
		m += uint(Pow(10, nb-1-i)) * d
		n = (n - d) / 10
	}

	return m
}

func IsPalindrome(n uint) bool {
	return n == Mirror(n)
}

func LargestPalindrome(n uint) (m uint) {

	for a := uint(Pow(10, n-1)); a < uint(Pow(10, n)); a++ {
		for b := uint(Pow(10, n-1)); b < uint(Pow(10, n)); b++ {
			if a*b > m {
				if IsPalindrome(a * b) {
					m = a * b
				}
			}
		}
	}
	return
}

func Reverse(s []uint) []uint {

	l := len(s)

	r := make([]uint, l, l)

	for k := 0; k < l; k++ {
		r[l-k-1] = s[k]
	}
	return r
}

func IntToSlice(n uint) []uint {
	nb := NbDigits(n)
	var d uint
	sl := make([]uint, nb, nb)

	for i := uint(0); i <= nb-1; i++ {
		d = n % 10
		sl[nb-i-1] = d
		n = (n - d) / 10
	}

	return sl
}

func IsPalindrome2(n uint) (b bool) {
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

func LargestPalindrome2(n uint) (m uint) {

	for a := uint(Pow(10, n-1)); a < uint(Pow(10, n)); a++ {
		for b := uint(Pow(10, n-1)); b < uint(Pow(10, n)); b++ {
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

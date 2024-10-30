package main

import (
	"fmt"
	"math"
)

func is_prime(n uint) bool {
	if n == uint(0) {
		return false
	}
	if n <= uint(3) {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%uint(i) == 0 || n%(uint(i)+2) == 0 {
			return false
		}
	}
	return true
}

func find_in_increasing_arr(x int, a []int) (b bool) {
	c := 0
	d := len(a) - 1
	var e int
	for d-c > 2 {
		e = (c + d) / 2
		if a[e] > x {
			d = e
		} else {
			c = e
		}
	}
	for k := c; k <= d; k++ {
		if a[k] == x {
			b = true
			return
		}
	}
	return
}

func str_to_uint(s string) (n uint) {
	for i, d := len(s)-1, 1; i >= 0; i, d = i-1, d*10 {
		n += uint(int(s[i]-'0')) * uint(d)
	}
	return
} // Note : it should be possible to implement this using a "dictionnaty" ('0' char <-> 0 uint)

func avance(n string, m uint) {
	if is_prime(str_to_uint(n)) || n == "" {
		if uint(len(n)) == m && n[0] != '0' {
			fmt.Println(n)
		}
		avance("0"+n, m)
		avance("1"+n, m)
		avance("2"+n, m)
		avance("3"+n, m)
		avance("4"+n, m)
		avance("5"+n, m)
		avance("6"+n, m)
		avance("7"+n, m)
		avance("8"+n, m)
		avance("9"+n, m)
	}
}

func main() {
	// avance("", 3)
	fmt.Println(5)
	fmt.Println(Primes0[100])
}

package main

import (
	"fmt"
)

type tree struct {
	root  string
	zero  *tree
	one   *tree
	two   *tree
	three *tree
	four  *tree
	five  *tree
	six   *tree
	seven *tree
	eight *tree
	nine  *tree
}

/*
type tree struct {
	zero  tree
	one   tree
	two   tree
	three tree
	four  tree
	five  tree
	six   tree
	seven tree
	eight tree
	nine  tree
}
doesn't work because types can't refer to themselves.
Although, if a type refers to itself _through pointers_ (as above), then it works.
*/

func new_value(t *tree, n uint) *tree {
	if n == 0 {
		return &tree{
			t.root,
			&tree{root: "0" + t.root},
			&tree{root: "1" + t.root},
			&tree{root: "2" + t.root},
			&tree{root: "3" + t.root},
			&tree{root: "4" + t.root},
			&tree{root: "5" + t.root},
			&tree{root: "6" + t.root},
			&tree{root: "7" + t.root},
			&tree{root: "8" + t.root},
			&tree{root: "9" + t.root},
		}
	}
	return &tree{
		t.root,
		new_value(t.zero, n-1),
		new_value(t.one, n-1),
		new_value(t.two, n-1),
		new_value(t.three, n-1),
		new_value(t.four, n-1),
		new_value(t.five, n-1),
		new_value(t.six, n-1),
		new_value(t.seven, n-1),
		new_value(t.eight, n-1),
		new_value(t.nine, n-1),
	}
}

func nth_tree_full(n uint) (t *tree) {
	for k := uint(0); k <= n; k++ {
		t = new_value(t, k)
	}
	return
}

func pow(x int, n uint) int {
	if n == 0 {
		return 1
	}
	return x * pow(x, n-1)
}

func str_to_uint(s string) (x uint) {
	x = 0
	for i, m := 0, len(s); i <= m; i++ {
		x += uint(int(s[i]-'0') * pow(10, uint(i)))
	}
	return
}

func is_prime(n uint) (b bool) {
	b = false
	fmt.Println(n)
	return
}

func dfs(t *tree, n uint) {

}

var Arr [5]int = [5]int{1, 2, 3, 4, 5}

/*
int(math.Pow(float64(10), float64(3)))
Go doesn't interpret this as a constant so it does not work.

func pow(x int, n uint) int {
	if n == 0 {
		return 1
	}
	return x * pow(x, n-1)
}
const p int = pow(3,5)



//This does not work because

/*
func dfs(n uint) [p]int {
	var a [p]int
	a[0] = n
	return a
}
*/

/*
func nthTree_full(n uint) *tree {
	if n == 0 {
		return &tree{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	} else {
		var p *uint
		*p = 1
		return &tree{p, nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1)}
	}
}
*/

func print() {
	str := "1357924680"
	for i := 0; i <= 9; i++ {
		fmt.Println(int(str[i] - '0'))
	}
}

/*
func nthTree_full(n uint, root *uint) *tree {
	if n == 0 {
		return &tree{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	} else {
		var p *uint
		*p = 1
		return &tree{p, nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1), nthTree_full(n - 1)}
	}
}
*/

/*
func nonworking_nthTree_1(n uint) tree {
	if n == 0 {
		return tree{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	} else {
		return tree{&(nonworking_nthTree(n - 1)), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1), &nonworking_nthTree(n - 1)}
	}
}
-> This doesn't work because the r	eturn values nonworking_nthTree(n - 1) are temporary values on the stack;
and the Go language doesn't allow taking the address of temporary values directly to avoid memory-safety issues.

func nonworking_nthTree_1(n uint) tree {
	if n == 0 {
		return tree{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	} else {
	 	a := nonworking_nthTree(n - 1)
		b := nonworking_nthTree(n - 1)
		c := nonworking_nthTree(n - 1)
		d := nonworking_nthTree(n - 1)
		e := nonworking_nthTree(n - 1)
		f := nonworking_nthTree(n - 1)
		g := nonworking_nthTree(n - 1)
		h := nonworking_nthTree(n - 1)
		i := nonworking_nthTree(n - 1)
		return tree{a,b,c,d,e,f,g,h,i}
	}
}
This still doesn't work because the calls to nth_tree are still creating temporary values.
One must use pointers to have non-temporary values : so long as a pointer exists, some memory is allocated to it
(and henceforth this creates a non temporary variable)
*/

func main() {
	t := new(tree)
	*t = tree{
		root: "",
	}
	fmt.Println(*t)
	t = new_value(t, 0)
	fmt.Println(*t)
	fmt.Println(*t.five)

	t = new_value(t, 1)
	fmt.Println(*t)
	fmt.Println(t.five)
	fmt.Println(t.five.five)
}

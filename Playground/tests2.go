package main

import (
	"encoding/json"
	"fmt"
)

type PreResponseWithRank struct {
	p int
	s float64
	v int
}

type Test struct {
	p string
	s string
	v string
}

func main() {

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
}

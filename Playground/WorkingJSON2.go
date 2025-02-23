package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	byt := []byte(`{"num":6.13,"str":"cakah","int":45}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Fatal(err)
	}

	{
		elem, ok := dat["int"].(float64) // Unmarshal always interprets numerical values as floats !
		if ok {
			fmt.Println(elem)
		} else {
			log.Fatal()
		}
	}

	{
		elem, ok := dat["num"].(float64)
		if ok {
			fmt.Println(elem)
		} else {
			log.Fatal()
		}
	}

	{
		elem, ok := dat["str"].(string)
		if ok {
			fmt.Println(elem)
		} else {
			log.Fatal()
		}
	}

}

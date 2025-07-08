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

	// preresprank := &Test{
	// 	p: "h",
	// 	s: "h",
	// 	v: "h",
	// }

	var sl []map[string]interface{}

	if err := json.Unmarshal([]byte(`[{"p":"a","s":"b","v":"c"}]`), &sl); err != nil {
		log.Fatal(err)
	}

	// if err := json.Unmarshal([]byte(`{"p":214,"s":0.1799,"v":14124}`), preresprank); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("got here")
	fmt.Println(sl)

	// Get json.Unmarshal working (https://pkg.go.dev/encoding/json) and then come back to the project.
	// Make minimal examples like this to help.
}

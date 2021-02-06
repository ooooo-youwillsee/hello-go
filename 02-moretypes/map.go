package main

import (
	"fmt"
)

type Vertex struct {
	x , y int
}

func main () {
	// m :=map[string]Vertex{}
	var m = map[string]Vertex{}
	m["1"] = Vertex{1, 2}
	m["2"] = Vertex{3, 4}
	fmt.Println(m)


	a := map[string]Vertex{
		"1": {0,1},
		"2": {2,3},
	}
	for k,v := range a {
		fmt.Println(k, v)
	}

	b, ok := a["1"]
	if ok {
		fmt.Printf("b = %d \n", b)
	}
}
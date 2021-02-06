package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1,2})

	v := Vertex{1,2}
	v.x = 4
	fmt.Println(v)

	a := Vertex{2,3}
	p := &a
	p.x = 0
	fmt.Println(*p)

	b := Vertex{x:9}
	fmt.Println(b)
}
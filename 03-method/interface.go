package main

import "fmt"

type I interface {
	M()
}

type X struct {
	s string
}

func (x X) M() {
	fmt.Println(x.s)
}

func main() {

	var i I = X{"111"}
	i.M()
}
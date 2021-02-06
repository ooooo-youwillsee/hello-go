package main

import "fmt"

type A func(int, int) int

func compute(fn A,x int, y int) int {
	return fn(x, y)
}

func main() {
	funcA := func (x int , y int )  int {
		return x+y
	}
	fmt.Println(compute(funcA, 1, 2))
}
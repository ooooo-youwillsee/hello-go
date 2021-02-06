package main

import "fmt"


func main() {
	a := [3]int{1,2,3}
	fmt.Println(a)


	b := [5]int{}
	fmt.Println(b)

	c := a[1:3]
	fmt.Println(c)
	c[0] = 9
	fmt.Println(c)
	// A slice does not store any data, it just describes a section of an underlying array.
	fmt.Println(a)

	d := []struct {
		x int
		y bool
	}{
		{1, false},
		{2, true},
	}
	fmt.Println(d)
	
	e := []int{1,2,3,4,5,6,7,8,9,10,}
	for i,v := range e {
		fmt.Println(i, v)
	}
}
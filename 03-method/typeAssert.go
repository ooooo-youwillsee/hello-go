package main

import "fmt"

func main() {
	var i interface{} = "hello"

	fmt.Println(i)

	s, ok := i.(string)
	fmt.Println(s, ok)
}
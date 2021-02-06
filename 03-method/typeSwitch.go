package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case string: 
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	default:
		fmt.Println("other type", v)
	}
}

func main() {

	do(1)
	do("s")
	do(1.0)
}
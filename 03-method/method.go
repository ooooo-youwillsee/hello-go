package main

import (
	"fmt"
	"math"
)

type MyFloat64 float64

func (f MyFloat64) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	a := MyFloat64(-math.Sqrt(2))
	fmt.Println(a.abs())
}
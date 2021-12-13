package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)

	go func() {
		fmt.Println("1")
		c <- 1
		fmt.Println("2")
		c <- 2
		fmt.Println("3")
		c <- 3
		fmt.Println("never execute here")
	}()
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	time.Sleep(10 * time.Second)
}

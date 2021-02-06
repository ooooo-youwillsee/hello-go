package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {

	// go 和 defer 后面接一个语句
	go say("hello")
	go fmt.Println("111111")
	say("world")
}

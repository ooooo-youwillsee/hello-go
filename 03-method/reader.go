package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := strings.NewReader("hello go!!!")

	buf := make([]byte, 3)
	for {
		len, err := s.Read(buf)
		fmt.Printf("len = %d, val = %q \n", len, buf[:len])
		// 说明流到末尾了
		if err == io.EOF {
			fmt.Println("eof ~~")
			break
		}
	}
}
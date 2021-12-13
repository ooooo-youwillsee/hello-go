package hello

import (
	"errors"
	"fmt"
)

func Hello(str string) (string, error) {
	if str == "" {
		return "", errors.New("error !!!")
	}
	println("hello " + str)
	s := []string{"0", "1"}
	test(s)
	fmt.Printf("s[1]: %s\n", s[1])
	_ = make(chan int)
	return str, nil
}

func test(s []string) {
	s[1] = "0"
}

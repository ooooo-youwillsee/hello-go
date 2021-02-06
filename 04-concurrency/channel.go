package main

import "fmt"

func sum(nums []int, c chan int) {
	ans := 0
	for _, v := range nums {
		ans += v
	}
	fmt.Println("start", nums)
	c <- ans
	fmt.Println("end", nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := make(chan int)
	mid := len(nums) / 2
	go sum(nums[:mid], c)
	go sum(nums[mid:], c)
	a, b := <-c, <-c

	fmt.Printf("sum = %d\n", a+b)
}

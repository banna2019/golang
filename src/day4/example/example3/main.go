package main

import (
	"fmt"
	"time"
)

func recusive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	recusive(n + 1)
}

func factor(n int) int{
	if n == 1 {
		return 1
	}
	return factor(n - 1) * n
}

func fab(n int) int{
	if n <= 1{
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	// recusive(0)
	// fmt.Println(factor(5))
	for i := 0; i < 10 ;i++{
		fmt.Println(fab(i))
	}
}

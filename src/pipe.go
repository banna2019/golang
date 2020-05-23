package main

import (
	"fmt"
)

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	fmt.Println(len(pipe))

}

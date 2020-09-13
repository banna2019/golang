package main

import (
	"fmt"
	"goroute_example/goroute"
)

func main() {

	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 200, pipe) //这里如果需要起一个线程直接使用"go"即可

	sum := <-pipe
	fmt.Println("sum=", sum)

}

package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("hello world")
// }

// func add(a int, b int) int {
// 	var sum int
// 	sum = a + b

// 	return sum
// }

// func main() {
// 	// time.Sleep(10 * time.Second)

// 	//这里是单行注释
// 	// var c int
// 	// c = add(100, 200)
// 	// go test_goroute(300, 300)
// 	// fmt.Println("add(100,200)=", c)
// 	// fmt.Println("Hello World")

// 	//以下是多行注释
// 	/*
// 		for i := 0; i < 100; i++ {
// 			go test_print(i)
// 		}
// 		time.Sleep(10 * time.Second)
// 	*/

// 	// test_pipe()

// 	/* 这里代码会被死锁,锁住没法退出*/
// 	/*fmt.Println("start goroute")
// 	go test_pipe()
// 	fmt.Println("end goroute")
// 	time.Sleep(10 * time.Second)*/

// }

/*第一种写法*/

// var pipe chan int

// func add(a int, b int) int {
// 	var sum int
// 	sum = a + b

// 	time.Sleep(3 * time.Second)
// 	pipe <- sum
// 	return sum //这里return不能注释
// }

// func main() {

// 	pipe = make(chan int, 1)
// 	go add(2, 5)

// 	sum := <-pipe
// 	fmt.Println("sum=", sum)

// }

/*第二种写法*/

// func add(a int, b int, c chan int) int {
// 	var sum int
// 	sum = a + b

// 	time.Sleep(3 * time.Second)
// 	c <- sum
// 	return sum //这里return不能注释
// }

// func main() {
// 	var pipe chan int
// 	pipe = make(chan int, 1)
// 	go add(2, 5, pipe)

// 	sum := <-pipe
// 	fmt.Println("sum=", sum)

// }

/*go的多返回值*/
func main() {
	// sum, avg := calc(100, 200)
	// fmt.Println("sum=", sum, "avg=", avg)

	sum, _ := calc(100, 200) //这里的下划线"_"是将返回值忽略掉
	fmt.Println("sum=", sum)
}

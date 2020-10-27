package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	1.当主进程执行的比协程快时,主进程退出执行后,协程也会被退出
	2.当协程执行比主进程快时,协程正常退出;不影响主进程正常运行
*/

/*
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go test() //表示开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	time.Sleep(time.Second * 1)
}
*/

var wg sync.WaitGroup

func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("test() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程的计数器减1
}

func test1() {
	for i := 1; i < 10; i++ {
		fmt.Println("test1() 你好golang-", i)
		time.Sleep(time.Millisecond * 30)
	}
	wg.Done()
}

func main() {
	wg.Add(1) ////协程计数器加1
	go test() //表示开启一个协程
	wg.Add(1)
	go test1()
	for i := 1; i < 10; i++ {
		fmt.Println("main() 你好golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait() //等待协程执行完毕
	fmt.Println("主线程退出")
}

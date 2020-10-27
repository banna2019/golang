package main

import (
	"fmt"
	"sync"
	"time"
)

//go build -race main.go  编译后运行查看("-race": 表示程序运行的时候,有没有竞争关系)

/*
//未添加mutex处理,具有锁占用程序
var count = 0
var wg sync.WaitGroup

func test() {
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Millisecond)
	wg.Done()
}

func main() {
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
*/

//处理了锁占用的处理程序段
var count = 0
var wg sync.WaitGroup

var mutex sync.Mutex

func test() {
	mutex.Lock()
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()

}

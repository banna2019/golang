package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	需求: 使用goroutine和channel协同工作案例
	1.开启一个WriteDate的协程个向管道inChan中写入100条数据
	2.开启一个ReadData的协程读取inChan中写入的数据
	3.注意: WriteData和ReadData同时操作一个管道
	4.主线程必须等待操作完成后才可以退出

	goroutine结合Channel使用的简单demo,定义两个方法,一个方法给管理里面写数据,一个给管道里面读取数据
*/

var wg sync.WaitGroup

//写数据(关闭channel需要添加到,写入数据中)
func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("[读取]数据%v\n", v)
		time.Sleep(time.Millisecond * 60)
	}
	wg.Done()
}

func main() {

	var ch = make(chan int, 10)

	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()
	fmt.Println("退出......")
}
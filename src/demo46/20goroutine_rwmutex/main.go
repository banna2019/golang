package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

//写的方法(在加加了读写互斥锁之后,写是互斥的;只能有一个写入)
func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

//读的方法(在加加了读写互斥锁之后,读的话是并行操作)
func read() {
	mutex.RLock()
	fmt.Println("---执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

func main() {

	//开启10个协程执行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 开启10个协程执行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[4] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
}

/*读写锁*/
/*
func testRWMutex() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32

	a[8] = 10
	a[3] = 10
	a[4] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				time.Sleep(time.Millisecond)
				// fmt.Println(a)
				// b[8] = rand.Intn(100)
				rwLock.RUnlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count)) //因为这里也是共享操作,需要使用原子操作读取出数据来
}*/

/*互斥锁*/
func testRWMutex() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32

	a[8] = 10
	a[3] = 10
	a[4] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				lock.Lock()
				time.Sleep(time.Millisecond)
				// fmt.Println(a)
				// b[8] = rand.Intn(100)
				lock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count)) //因为这里也是共享操作,需要使用原子操作读取出数据来
}

func main() {
	// testMap()
	testRWMutex()
}

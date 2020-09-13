package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Millisecond * 1000)
}

func main() {

	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05")) //打印日期需要"2006/01/02 15:04:05",这个串不然是无法格式化的

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("const:%d\n", (end-start)/1000)

}

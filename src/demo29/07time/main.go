package main

import (
	"fmt"
	"time"
)

//日期字符串转换成时间戳
func main() {
	/*
		golang定时器
	*/

	/*
		//版本一
		// ticker := time.NewTicker(time.Second) //定义一个1秒间隔的定时器
		// n := 0
		// for i := range ticker.C {
		// 	fmt.Println(i) //每秒都是执行的任务
		// 	n++
		// 	if n > 5 {
		// 		ticker.Stop()
		// 		return
		// 	}
		// }

		//版本二
		ticker := time.NewTicker(time.Second) //定义一个1秒间隔的定时器
		n := 5
		for t := range ticker.C {
			n--
			fmt.Println(t)
			if n == 0 {
				ticker.Stop() //终止这个定时器继续执行
				break
			}
		}
	*/

	//休眠方法
	// fmt.Println("aaa")
	// time.Sleep(time.Second)
	// fmt.Println("aaa1")
	// time.Sleep(time.Second)
	// fmt.Println("aaa2")
	// time.Sleep(time.Second * 5)
	// fmt.Println("aaa3")

	for {
		time.Sleep(time.Second)
		fmt.Println("我在执行定时任务")
	}
}

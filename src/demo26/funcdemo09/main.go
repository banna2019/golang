package main

import "fmt"

func main() {
	/*
	   Go语言中函数内部并不能再像之前那样定义函数了,只能定义匿名函数.匿名函数就是没有函数名的函数
	*/

	//匿名函数 匿名字执行函数
	func() {
		fmt.Println("test......")
	}() //"()"执行此函数

	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(2, 3))

	//匿名自执行函数接收
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

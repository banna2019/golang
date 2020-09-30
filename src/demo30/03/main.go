package main

import "fmt"

func fn1(x int) {
	x = 10
}

func fn2(x *int) { //*int为引用类型
	*x = 40
}

func main() {
	var a = 5
	fn1(a)
	fmt.Println(a)

	fn2(&a) //&a为引用a的内存地址
	fmt.Println(a)
}

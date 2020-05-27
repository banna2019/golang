package main

import (
	"fmt"
)

func modify(a int) {
	a = 10
	return
}

func modify1(a *int) {
	*a = 10 //取实际的内存地址
}

func main() {
	a := 5
	b := make(chan int, 1)

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	modify(a)
	fmt.Println("a=", a)
	modify1(&a) //这里的"&a"使用的是内存地址,最终的值是内存地址上存储的值
	fmt.Println("a=", a)
}

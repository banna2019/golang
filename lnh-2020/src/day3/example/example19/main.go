package main

import "fmt"

func main() {
	var i int = 0
	defer fmt.Println(i)        //第二先出栈
	defer fmt.Println("second") //第一个先出栈

	i = 10
	fmt.Println(i) //本函数中最先输出的内容
}

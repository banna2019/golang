package main

import (
	"fmt"
)

func main() {
	var str = "hello world\n\n"
	var srt1 = `hello \n \n \n
	this is a test string
	This is a test string too.`

	var b byte = 'c'

	fmt.Println("str=", str)
	fmt.Println("str1=", srt1)
	fmt.Println(b) //字符输出的是一个数字
	fmt.Printf("%c\n", b)
}

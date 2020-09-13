package main

import "fmt"

func main() {
	/*
		fmt.Println("go", "python", "php", "javascript")
		fmt.Print("go", "python", "php", "javascript")
	*/

	/*
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	*/

	/*
		var a = "aaa" //go语言中变量中定义会后必须要使用

		fmt.Println(a)

		fmt.Printf("%v", a)
	*/

	a := 10
	b := 20
	c := 30

	fmt.Println("a=", a, ",b=", b, ",c=", c)
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)

	//时候用Printf打印一个变量的类型
	fmt.Printf("a=%v a的类型%T", a, a)

}

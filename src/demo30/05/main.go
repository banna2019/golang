package main

import "fmt"

func main() {

	//实际开发中,new函数不太常用,使用new函数得到的是一个指针,并且该指针对应用的值为该类型的零值
	var a = new(int)                                   //a是一个指针类型	类型是 *int的指针类型 值是0
	fmt.Printf("值: %v 类型:%T 指针变量对应的值: %v\n", a, a, *a) //值: 0xc0000b4008 类型:*int 指针变量对应的值: 0

	/*
		//错误写法
		var a *int //指针也是引用数据类型
		*a = 100
		fmt.Println(*a)
	*/

	//new方法给指针变量分配存储空间
	// var b *int
	// b = new(int)
	// *b = 100
	// fmt.Println(b)

	var f = new(bool)
	fmt.Println(*f)

	/*
		new 与make 的区别

			1.二者都是用来做内存分配的.

			2.make 只用于slice、map 以及channel 的初始化,返回的还是这三个引用类型本身

			3. 而new 用于类型的内存分配,并且内存对应的值为类型零值,返回的是指向类型的指针.
	*/
}

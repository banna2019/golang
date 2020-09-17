package main

import "fmt"

//golang 闭包
/*
全局变量特点:
	1.常驻内存
	2.污染全局

局部变量的特点;
	1.不常住内存
	2.不污染全局

闭包:
	1.可以让一个变量常驻内存
	2.可以让一个变量不污染全局

	注释: 闭包可以理解成"定义在一个函数内部的函数".在本质上,闭包是将函数内部和函数外部连接起来的桥梁.或者说是函数和其他引用环境的组合体.
*/

/*
闭包
	1.闭包是指有权访问另一个函数作用域中的变量的函数.
	2.创建闭包的常见的方式就是在一个函数内部创建另一个函数,通过另一个函数访问这个函数的局部变量

注意:
	由于闭包里作用域返回的是局部变量资源不会被立刻销毁回收,所以可能会占用更多的内存.
	过度使用闭包会导致性能下降,建议在非常有必要的时候才使用闭包.
*/

//写法: 闭包写法 函数里面嵌套一个函数 最后返回里面的函数
func adder() func() int {
	var i = 10 //常驻内存,不污染全局
	return func() int {
		return i + 1
	}
}

func adder1() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

// func adder1() func(int) int {
// 	var i = 10
// 	return func(y int) int {
// 		i += y
// 		return i
// 	}
// }

func main() {
	var fn = adder() //"()"表示执行方法
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	var fn1 = adder1()
	fmt.Println(fn1(10)) //20
	fmt.Println(fn1(10)) //30
	fmt.Println(fn1(10)) //40
}

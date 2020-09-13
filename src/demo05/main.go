package main

import "fmt"

func main() {

	// var username = "张三"
	// username = "李四"
	// username = "王五"

	// fmt.Println(username)

	//常量(变量的值可以改变,常量的值不可改变)
	//常量在定义的时候必须要赋值

	// const pi = 3.14159
	// fmt.Println(pi)

	/*
		//错误写法
		const pi string
		pi="xxxx"
		fmt.Println(pi)
	*/

	// const A = "A"
	// A = "AAA"	//cannot assign to A 常量的值不可改变
	// fmt.Println(A)

	/*
		//多个常量也可以一起声明
		const (
			A = "A"
			B = "B"
		)

		fmt.Println(A, B)
	*/

	/*
		//const 同时声明多个常量时,如果省略了值则表示和上面一行的值相同
		const (
			n1 = 100
			n2
			n3
		)

		fmt.Println(n1, n2, n3)
	*/

	/*
		iota 是golang语言的常量计数器,只能在常量的表达式中使用.
		iota 在const关键字出现时将被重置为0(const内部的第一行之前),const中每新增一行常量声明将使iota 计数一次(iota 可理解为const 语句块中的行索引).
	*/
	// const a = iota
	// fmt.Println(a)

	// const (
	// 	n1 = iota
	// 	n2
	// 	n3
	// 	n4
	// )

	// fmt.Println(n1, n2, n3, n4)

	// const (
	// 	n1 = iota
	// 	_
	// 	n3
	// 	n4
	// )

	// fmt.Println(n1, n3, n4)

	//iota插队
	// const (
	// 	n1 = iota //0
	// 	n2 = 100  //100
	// 	n3 = iota //2
	// 	n4        //3
	// )

	// fmt.Println(n1, n2, n3, n4)

	//多个iota定义在一行
	// const (
	// 	n1, n2 = iota + 1, iota + 2
	// 	n3, n4
	// 	n5, n6
	// )

	// fmt.Println(n1, n2, n3, n4, n5, n6)

	//go语言中变量的名字是区分大小写的
	var age = 20
	var Age = 30

	fmt.Println(age, Age)
}

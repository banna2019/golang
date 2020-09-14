//逻辑运算
package main

import "fmt"

//定义一个方法
func test() bool {

	fmt.Println("test...")
	return true
}

func main() {
	/*
	   &&	逻辑 AND 运算符.如果两边的操作数都是true,则为true,否则为false
	   ||	逻辑 OR 运算符.如果两边的操作数有一个true,则为true,否则为false
	   !	逻辑 NOT 运算符.如果条件为true,则为false,否则为true
	*/

	// var a = 23
	// var b = 8

	// fmt.Println(a > 10 && b < 10) //true
	// fmt.Println(a > 24 && b < 10) //false
	// fmt.Println(a > 5 && b < 6)   //false
	// fmt.Println(a == 5 && b < 6)   //false

	// var a = 23
	// var b = 8

	// fmt.Println(a > 10 || b < 10) //true
	// fmt.Println(a > 24 || b < 10) //true
	// fmt.Println(a > 5 || b < 6)   //true
	// fmt.Println(a == 5 || b < 6)  //false

	// flag := false
	// fmt.Println(!flag)

	//逻辑与和逻辑或短路
	//逻辑与: 前面是false后面就不会执行	逻辑或: 前面是true后面就不会执行
	/*
		输出:
			test...
			执行
	*/
	//逻辑与
	// var a = 10
	// // if a > 9 && test() {
	// if a > 11 && test() {
	// 	fmt.Println("执行")
	// }

	/*
		输出(空):

	*/
	//逻辑与
	// var a = 10
	// if a > 11 && test() {
	// 	fmt.Println("执行")
	// }

	/*
		输出:
			test...
			执行
	*/
	//逻辑或
	// var a = 10
	// if a > 11 || test() {
	// 	fmt.Println("执行")
	// }

	/*
		输出:
			执行
	*/
	//逻辑或
	var a = 10
	if a < 11 || test() {
		fmt.Println("执行")
	}
}

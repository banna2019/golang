package main

import (
	"fmt"
)

/*函数可以赋值给一个变量*/
// func add(a, b int) int {
// 	return a + b
// }

// func main() {
// 	c := add //函数可以赋值给一个变量
// 	sum := c(200, 300)
// 	fmt.Println(sum)
// }

/*自定义一种函数类型一*/
// type op_func func(int, int) int //函数也是一种类型

// func add(a, b int) int {
// 	return a + b
// }

// func operator(op op_func, a, b int) int {
// 	return op(a, b)
// }

// func main() {
// 	c := add
// 	sum := operator(c, 100, 200)
// 	fmt.Println(sum)
// }

/*自定义一种函数类型二*/
// func add(a, b int) int {
// 	return a + b
// }

// func operator(op func(int, int) int, a, b int) int { //这里func里面有时候会返回很多值,这样写会比较繁琐
// 	return op(a, b)
// }

// func main() {
// 	c := add
// 	sum := operator(c, 100, 200)
// 	fmt.Println(sum)
// }

type op_func func(int, int) int

func sub(a, b int) int { //这里的"func sub(a, b int) int"函数没有特殊声明类型,及函数的返回值都同一种类型;这是为同一种类型的函数
	return a - b
}

func add(a, b int) int {
	return a + b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

/*写法一*/
// func main() {
// 	sum := operator(sub, 100, 200)
// 	fmt.Println(sum)
// }

/*写法二*/
func main() {
	var c op_func
	c = add
	sum := operator(c, 100, 200)
	fmt.Println(sum)

	fmt.Println(c)
	fmt.Println(add)
}

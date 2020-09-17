package main

import "fmt"

//函数作为另一个函数参数(方法作为另一个方法的参数)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//自定义一个方法类型
type calcType func(int, int) int

func calc(x, y int, cb calcType) int { //这类方法作为另一个方法的参数"func(int,int)int)"
	return cb(x, y)
}

func main() {

	// sum := calc(10, 5, add)
	// fmt.Println(sum)

	// s := calc(10, 5, sub)
	// fmt.Println(s)

	j := calc(3, 4, func(x, y int) int { //这里的"func(x, y int) int"为匿名函数
		return x * y
	})
	fmt.Println(j)
}

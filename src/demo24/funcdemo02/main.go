package main

import "fmt"

func sumFn(x, y int) int {
	return x + y
}

//return 关键词一次可以返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名: 函数定义时可以给返回值命名,并在函数体中直接使用这些变量,最后通过return关键字
func calc1(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub) //开始没有被初始化为"0"
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub)
	return
}

func calc2(x, y int) (sum, sub int) { //返回值参数的简写
	fmt.Println(sum, sub) //开始没有被初始化为"0"
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub)
	return
}

func main() {
	// sum1 := sumFn(10, 2)
	// fmt.Println(sum1)

	// a, b := calc(10, 2)
	// fmt.Println(a, b)

	// a, b := calc1(10, 2)
	// fmt.Println(a, "-", b)

	// a, b := calc2(10, 2)
	// fmt.Println(a, "-", b)

	a, _ := calc2(10, 2) //"_"为匿名变量可以被忽略
	fmt.Println(a)
}

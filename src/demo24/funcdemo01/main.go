package main

import "fmt"

//求两个数的和
func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

//求两个数的差
func subFn(x int, y int) int { //这里的"(x int, y int)"为形参
	sub := x - y
	return sub
}

//函数参数的简写
func subFn1(x, y int) int { //这里的"(x, y int)"为形参
	sub := x - y
	return sub
}

//函数的可变参数,可变参数是指函数的参数数量不固定.Go 语言中可变参数通过在参数后加...来标识
func sumFn1(x ...int) int {

	// fmt.Printf("%v - %T", x, x) //[12 34 45 46] - []int

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn2(x int, y ...int) int { //"...int"在类型前面跟3个点,这就是一个切片
	fmt.Println(x, y)

	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func main() {
	// sum1 := sumFn(12, 3)
	// fmt.Println(sum1)

	// sum2 := sumFn(15, 5)
	// fmt.Println(sum2)

	// a := 20
	// b := 2
	// sub1 := subFn(a, b) //这里的"(a, b)"为实参
	// fmt.Println(sub1)

	// sub2 := subFn1(22, 4)
	// fmt.Println(sub2)

	// sum3 := sumFn1(12, 34, 45, 46)
	// fmt.Println(sum3)

	// sum4 := sumFn1(120, 34, 455, 436)
	// fmt.Println(sum4)

	sum5 := sumFn2(100, 1, 2, 4, 6, 7)
	fmt.Println(sum5)
}
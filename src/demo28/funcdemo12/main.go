package main

import "fmt"

/*
	Go语言中的defer语句会将其后面跟随的语句进行延迟处理.
	在defer归属的函数即将返回时,将延迟处理的语句按defer定义的逆序进行执行,
	也就是说,先被defer的语句最后被执行,最后被defer的语句,最先被执行.
*/

/*
func f1() {
	fmt.Println("开始")

	defer func() { //如果要使用defer来执行一系列延迟操作,那么就需要把这些操作放到匿名自执行函数中
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()
	fmt.Println("结束")
}

//调用方法返回的是0	fmt.Println(f2()) //0
func f2() int {
	var a int

	defer func() {
		a++ //匿名返回值,操作以前的变量不受影响
	}()
	// fmt.Println("结束")
	return a
}

func f3() (a int) {	//命名返回值是操作以后的值
	defer func() {
		a++
	}()
	return a
}
*/

func main() {
	//1.defer的使用演示
	//如果程序中有多个defer定义语句,在执行defer语句的时候,是逆序执行的
	// fmt.Println("开始")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("结束")

	//2.defer在命名返回值和匿名返回 函数表现不一样
	// f1()
	fmt.Println(f2()) //0
	fmt.Println(f3()) //1

}

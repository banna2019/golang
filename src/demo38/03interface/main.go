package main

import "fmt"

/*
	Golang中的接口可以不定义任何方法,没有定义任何方法的接口就是空接口.空接口表示没有任何约束,因此任何类型变量都可以实现空接口.
	空接口在实际项目中用的是非常多的,用空接口可以表示任意数据类型.
*/

type A interface{} //空接口	表示没有任何的约束	任意的类型都可以实现空接口

func main() {

	var a A
	var str = "你好golang"
	a = str //让字符串实现A这个接口
	fmt.Printf("值:%v 类型:%T\n", a, a)

	var num = 20
	a = num //表示让int类型实现A这个接口
	fmt.Printf("值: %v 类型:%T\n", a, a)

	var flag = true
	a = flag //表示让bool类型实现A这个接口
	fmt.Printf("值: %v 类型:%T", a, a)
}

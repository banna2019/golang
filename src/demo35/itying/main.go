package main

import (
	"fmt"
	"itying/calc"
	T "itying/tools" //引入包的别名设置
	// _ "itying/tools" //表示匿名引入这个包
)

//main包中init函数 优先于 main函数
func init() { //这个方法是被系统调用的
	fmt.Println("main init.....")
}

func main() {

	sum := calc.Add(10, 2)
	fmt.Println(sum)
	fmt.Println(calc.Age)
	// fmt.Println(calc.aaa)

	sub := calc.Sub(10, 2)
	fmt.Println(sub)

	//调用tools包里面的方法
	// b := tools.Mul(2, 3)
	// fmt.Println(b)

	// tools.PrintInfo()
	// tools.SortIntAsc()

	T.Mul(3, 4)
	T.PrintInfo()
	T.SortIntAsc()
}

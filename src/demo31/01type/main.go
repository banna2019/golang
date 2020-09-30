package main

import "fmt"

//自定义类型
type myInt int

// type myFn func(int,int) int

//类型别名
type myFloat = float64

func main() {

	var a myInt = 10
	fmt.Printf("%v %T\n", a, a) //自定义类型,这里输出的类型就是;自定义的类型

	var b myFloat = 12.3
	fmt.Printf("%v %T", b, b) //类型别名这里输出的类型还是之前的类型
}

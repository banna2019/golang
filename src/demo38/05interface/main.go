package main

import "fmt"

/*
	类型断言
		一个接口的值(简称接口值)是由一个具体类型和具体类型的值两部分组成的.这两部分分别称为接口的动态类型和动态值.
		如果想要判断空中值的类型,那么这个时候就可以使用类型断言.
*/

//定义一个方法,可以传入任意数据类型,然后根据不同的类型实现不同的功能
func MyPrint(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("int类型")
	} else if _, ok := x.(bool); ok {
		fmt.Println("bool类型")
	}
}

//x.(type)判断一个变量的类型 这个语句只能用在switch语句里
func MyPrint1(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("传入错误")
	}
}

func main() {
	// var a interface{}
	// a = "你好golang"
	// v, ok := a.(string)
	// if ok {
	// 	fmt.Println("a就是一个string类型,值是: ", v)
	// } else {
	// 	fmt.Println("断言失败")
	// }

	MyPrint("你好golang")
	MyPrint(true)

	MyPrint1("你好golang")
	MyPrint1(true)
}

package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	// fmt.Println(x)

	//var num = 10 + x //invalid operation: 10 + x (mismatched types int and interface {});这里的x是空接口类型
	/*
		//使用类型断言进行处理
		b, _ := x.(int)
		var num = 10 + b
		fmt.Println(num)
	*/

	/*
		//反射来实现这个功能
		v := reflect.ValueOf(x)
		fmt.Println(v)

		var n = v + 12 //invalid operation: v + 12 (mismatched types reflect.Value and int)
		fmt.Println(n)
	*/

	//反射获取变量的原始值
	v := reflect.ValueOf(x)
	var m = v.Int() + 12
	fmt.Println(m)

}

func main() {

	var a = 13
	reflectValue(a)
}

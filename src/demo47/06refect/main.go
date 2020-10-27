package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}) {
	// *x = 120 //invalid indirect of x (type interface {})

	/*
		v, _ := x.(*int)
		*v = 120 //invalid memory address or nil pointer dereference
	*/

	v := reflect.ValueOf(x)
	// fmt.Println(v.Kind()) //ptr(指针类型)

	// fmt.Println(v.Elem().Kind()) //int64(使用"Elem()"获取类型)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("golang")
	}
}

func main() {
	var a int64 = 100
	reflectSetValue(&a)
	fmt.Println(a)

	var b string = "你好golang"
	reflectSetValue(&b)
	fmt.Println(b)
}

package main

import "fmt"

//golang中是允许给,结构体和自定义类型来自定义方法的
//注意事项: 非本地类型并不能定义方法,也就是说并不能给别的包的类型定义方法.
type MyInt int

func (m MyInt) PrintInfo() {
	fmt.Println("我是自定义类型里面的自定义方法")
}

func main() {
	var a MyInt = 20
	a.PrintInfo()

}

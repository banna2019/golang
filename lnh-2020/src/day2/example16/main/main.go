package main

import "fmt"

func main() {
	var a int = 100
	var b bool
	var c byte = 'a'
	d := 'b'

	fmt.Printf("%v\n", a) //"%v"原样输出结果
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("90%%\n")
	fmt.Printf("%t\n", b)
	fmt.Printf("%b\n", 100) //二进制方式打印
	fmt.Printf("%e\n", 100000000)
	fmt.Printf("%f\n", 199.22)
	fmt.Printf("%q\n", "this is a test")
	fmt.Printf("%x\n", 39839333) //十六进制
	fmt.Printf("%p\n", &a)       //取地址

	str := fmt.Sprintf("%d", a)
	fmt.Printf("%q", str)
}

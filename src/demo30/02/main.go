package main

import "fmt"

func main() {
	a := 10

	p := &a //p是指针变量  类型 *int

	// *p:	表示取出p这个变量对应的内存地址的值

	fmt.Println(&a) //输出的是a的地址
	fmt.Println(p)  //输出的是a的地址
	fmt.Println(*p) //*p 表示打印a对应的值

	*p = 30 //改变p这个变量对应的内存地址的值

	fmt.Println(a)
}

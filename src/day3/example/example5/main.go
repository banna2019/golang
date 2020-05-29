package main

import "fmt"

func modify(p *int) {
	fmt.Println(p)
	*p = 10000
	return
}

func main() {
	var a int = 10
	fmt.Println(&a) //打印a变量的地址

	var p *int //类型前面添加*就是,指针类型
	p = &a
	fmt.Println(*p) //要获取指针指向地址的值,直接在前面添加"*"即可

	*p = 100 //修改了指向地址存储的值
	fmt.Println(a)

	var b int = 999
	p = &b
	*p = 5
	fmt.Println(a) //这里a的值是不会发生变化的,这里的指针指向的是变量b
	fmt.Println(b)

	modify(&a) //这里需要取地址
	fmt.Println(a)
}

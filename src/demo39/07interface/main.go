package main

import "fmt"

/*
	结构体值接收者实现接口:
		值接收者: 如果结构体中的方法是值接收者,那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口类型变量
*/

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() { //值接收者
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {
	//结构体值接收者实例化后的结构体和结构体指针类型都可以赋值给接口变量

	var p = Phone{
		Name: "小米手机",
	}

	var p1 Usber = p //表示让Phone实现Usb的接口
	p1.start()

	var p2 = &Phone{
		Name: "小米手机",
	}

	var p3 Usber = p2 //表示让Phone实现Usb的接口
	p3.start()
}

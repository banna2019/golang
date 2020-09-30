package main

import "fmt"

/*
	在Golang中接口(interface)是一种类型,一种抽象的类型.接口(interface)是一组函数method的集合,Golang中的接口不能包含任何变量
	Golang中接口中的所有哦方法都没有方法体,接口定义了一个对象的行为规范,只定义规范不实现.接口体现了程序设计的多态和高内聚低耦合的思想
	Golang中的接口也是一种数据类型,不需要显示实现.只需要一个变量含有接口类型中的所有方法,那么这个变量就实现了这个接口.

	备注:
		谁实现了接口就需要实现这个接口中的方法
*/

//接口是一个规范
type Usber interface {
	start()
	stop()
}

//如果接口里面有方法的话,必须要通过结构体或者通过自定义类型实现这个接口

type Phone struct {
	Name string
}

//手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关机")
}

func (p Camera) run() {
	fmt.Println("run")
}

func main() {
	p := Phone{
		Name: "华为手机",
	}

	p.start()

	var p1 Usber //golang中接口就是一个数据类型

	p1 = p //表示手机实现Usb接口

	p1.start()
	p1.stop()

	c := Camera{} //实例化Camera

	var c1 Usber = c //表示相机实现了Usb接口
	c1.start()
	c1.stop()
	// c1.run() //错误: c1.run undefined (type Usber has no field or method run)
	c.run()
}

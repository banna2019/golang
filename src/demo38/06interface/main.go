package main

import "fmt"

/*
	在Golang中接口(interface)是一种类型,一种抽象的类型.接口(interface)是一组函数method的集合,Golang中的接口不能包含任何变量
	Golang中接口中的所有哦方法都没有方法体,接口定义了一个对象的行为规范,只定义规范不实现.接口体现了程序设计的多态和高内聚低耦合的思想
	Golang中的接口也是一种数据类型,不需要显示实现.只需要一个变量含有接口类型中的所有方法,那么这个变量就实现了这个接口.
*/

type Usber interface {
	start()
	stop()
}

//电脑结构体
type Computer struct {
}

func (c Computer) work(usb Usber) {
	//这个时候就需要判断usb的类型
	if _, ok := usb.(Phone); ok { //类型断言
		usb.start()
	} else {
		usb.stop()
	}
}

//手机结构体
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

//相机结构体
type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关闭")
}

func main() {

	var computer = Computer{}
	var phone = Phone{
		Name: "小米手机",
	}
	var camera = Camera{}

	computer.work(phone)
	computer.work(camera)
}

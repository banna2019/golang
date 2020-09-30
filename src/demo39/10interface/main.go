package main

import "fmt"

//定义一个Animal的接口,Animal中定义两个方法,分别是SetName和GetName.分别让Dog结构体和Cat结构体实现这个方法

type Animaler interface {
	SetName(string)
}

type Animaler1 interface {
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

//一次实现多个接口(一个结构体实现多个方法)
func main() {

	//Dog实现Animal的接口
	d := &Dog{
		Name: "小黑",
	}
	var d1 Animaler = d  //表示让Dog实现Animal这个接口
	var d2 Animaler1 = d //表示让Dog实现Animal1这个接口

	d1.SetName("小花狗")
	fmt.Println(d2.GetName())

}

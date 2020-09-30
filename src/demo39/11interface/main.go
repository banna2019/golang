package main

import "fmt"

type Ainterface interface {
	SetName(string)
}

type Binterface interface {
	GetName() string
}

type Animaler interface { //接口嵌套
	Ainterface
	Binterface
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
	var d1 Animaler = d //表示让Dog实现Animal这个接口
	d1.SetName("小花狗")
	fmt.Println(d1.GetName())

}

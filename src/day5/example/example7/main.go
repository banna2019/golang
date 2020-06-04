package main

import "fmt"

type integer int

func (p integer) print() {
	fmt.Println(p)
}

func (p *integer) set(b integer) {
	*p = b
}

type Student struct {
	Nmae  string
	Age   int
	Score int
}

func (p *Student) init(name string, age int, score int) {
	p.Nmae = name
	p.Age = age
	p.Score = score

	fmt.Println(p)
}

func (p Student) get() Student {
	return p
}

func main() {
	var stu Student
	// (&stu).init("stu01", 10, 200) //这种写法是原始传入引用地址

	stu.init("stu01", 10, 200) //这样写默认是会,推导出结构体
	stu1 := stu.get()
	fmt.Println(stu1)

	var a integer
	a = 100
	a.print()

	a.set(1000)
	a.print()
}

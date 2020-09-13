package main

import "fmt"

//接口里面可以包含多个方法;接口也是一种类型
type Test interface {
	Printf()
	Sleep()
}

type People struct {
	name   string
	age    int
	gender string
}

type Student struct {
	name  string
	age   int
	score int
}

//指针地址引用
/*
func (p *Student) Printf() { //1.这里的"*Student"是指针类型
	fmt.Println("name:", p.name)
	fmt.Println("age", p.age)
	fmt.Println("score:", p.score)
}

func main() {
	var t Test
	var stu Student = Student{
		name:  "stu01",
		age:   20,
		score: 100,
	}

	t = &stu //1.所以这里需要,指向stu的内存地址"&stu"
	t.Printf()
}*/

func (p Student) Printf() {
	fmt.Println("name:", p.name)
	fmt.Println("age", p.age)
	fmt.Println("score:", p.score)
}

func (p Student) Sleep() {
	fmt.Println("Student Sleep...")
}

func (p People) Printf() {
	fmt.Println("name:", p.name)
	fmt.Println("age", p.age)
	fmt.Println("gender:", p.gender)
}

func (p People) Sleep() {
	fmt.Println("People sleep...")
}

func main() {
	var t Test
	var stu Student = Student{
		name:  "stu01",
		age:   20,
		score: 100,
	}

	t = stu
	fmt.Println()
	t.Printf()
	t.Sleep()

	var people People = People{
		name:   "bruce",
		age:    30,
		gender: "male",
	}

	t = people
	fmt.Println()
	t.Printf()
	t.Sleep()

	/*
		备注:
			这里的"t"可以指向多种形态称之为多态.
	*/
}

package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {
	/*结构体初始化1*/
	var stu Student
	stu.Age = 18
	stu.Name = "xiaoming"
	stu.score = 99

	/*结构体初始化2*/
	var stu1 *Student = &Student{
		Age:  20,
		Name: "xiaoming-B",
	}

	fmt.Println(stu1.Name)

	/*结构体初始化3*/
	var stu2 Student = Student{
		Age:  21,
		Name: "xiaoming-C",
	}

	fmt.Println(stu2.Name)

	fmt.Println(stu)
	fmt.Printf("Name: %p\nAge: %p\nscore: %p\n", &stu.Name, &stu.Age, &stu.Age)
}

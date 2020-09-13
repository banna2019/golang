package main

import "fmt"

type interger int

type Student struct {
	Number int
}

type Stu Student

func main() {
	var i interger = 1000
	var j int = 100
	j = int(i)
	fmt.Println(i)
	fmt.Println(j)

	var a Student
	a = Student{30}

	var b Stu
	a = Student(b)
	fmt.Println(a)
}

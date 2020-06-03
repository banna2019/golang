package main

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func main() {
	var root *Student = new(Student)

	root.Name = "xiaoming"
	root.Age = 18
	root.Score = 100
	//一下两行代码默认为nil所以可以忽略不写
	// root.left = nil
	// root.right = nil

	var left1 *Student = new(Student)

	left1.Name = "xiaoming-left"
	left1.Age = 18
	left1.Score = 100
}

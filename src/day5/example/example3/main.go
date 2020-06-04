package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	//前序遍历
	// fmt.Println(root)
	// trans(root.left)
	// trans(root.right)

	//中序遍历
	// trans(root.left)
	// fmt.Println(root)
	// trans(root.right)

	//后序遍历
	trans(root.left)
	trans(root.right)
	fmt.Println(root)
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

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "xiaoming-right"
	right1.Age = 18
	right1.Score = 100

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "xiaoming-left1"
	left2.Age = 18
	left2.Score = 100

	left1.left = left2

	trans(root)
}

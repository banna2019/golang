package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		// p = (*p).next	//标准写法
		p = p.next //简化写法
	}
	fmt.Println()
}

/*func main() { //这里的链表使用的是尾部插入法
	var heard Student
	heard.Name = "xiaoming-A"
	heard.Age = 18
	heard.Score = 100

	var stu1 Student
	stu1.Name = "xiaoming-B"
	stu1.Age = 23
	stu1.Score = 110
	// stu1.next = nil	//结尾值默认是nil,所以可以不用写

	heard.next = &stu1 //链表串联
	trans(&heard)

	var stu2 Student
	stu2.Name = "xiaoming-C"
	stu2.Age = 24
	stu2.Score = 120

	stu1.next = &stu2 //添加链表,把stu1的next指向stu2的内存地址
	trans(&heard)

	var stu3 Student
	stu3.Name = "xiaoming-D"
	stu3.Age = 25
	stu3.Score = 130

	stu2.next = &stu3
	trans(&heard)
}*/

func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {
		stu := &Student{ //这里直接指向内存地址
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = stu //这里赋值即可
		tail = stu      //这里赋值即可
	}
}

/*for循环从尾部插入next,链表写法1*/
/*
func main() { //这里的链表使用的是尾部插入法
	var heard Student
	heard.Name = "xiaoming-A"
	heard.Age = 18
	heard.Score = 100

	var tail = &heard
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
	trans(&heard)
}*/

/*for循环从尾部插入next,链表写法2*/
/*
func main() { //这里的链表使用的是尾部插入法
	var heard Student
	heard.Name = "xiaoming-A"
	heard.Age = 18
	heard.Score = 100

	insertTail(&heard)
	trans(&heard)
}*/

/*for循环从头部插入next,链表写法1(头部插入式逆序的)*/
/*func main() {
	// var heard *Student = &Student{}
	var heard *Student = new(Student) //这里需要分配内存
	heard.Name = "xiaoming"
	heard.Age = 18
	heard.Score = 100

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		//问题段代码
		//stu.next = &heard
		// heard = stu ----->{heard.Name = stu.Name,heard.Age = stu.Age,heard.next = stu.next}
		//fmt.Println(stu)
		//time.Sleep(time.Second)

		stu.next = heard
		heard = &stu
	}

	trans(heard)
}*/

/*指针的指针的写法*/
func insertHeard(p **Student) { //这里是指针的指针
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		stu.next = *p
		*p = &stu //这里只是该变了指针副本的值,最终它外面的值是不会发生该变的
	}
}

func delNode(p *Student) { //这里删除节点会有问题,不能删除头节点;如果删除头结点,后面的节点没法串联
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6" {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}

func addNode(p *Student, newNode *Student) {
	for p != nil {
		if p.Name == "stu7" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}

func main() {
	// var heard *Student = &Student{}
	// var heard *Student = new(Student) //这里需要分配内存
	var heard *Student = new(Student) //这里需要分配内存
	heard.Name = "xiaoming"
	heard.Age = 18
	heard.Score = 100

	insertHeard(&heard)
	trans(heard)

	delNode(heard)
	trans(heard)

	var newNode *Student = new(Student) //这里需要分配内存
	newNode.Name = "stu1000"
	newNode.Age = 18
	newNode.Score = 100

	addNode(heard, newNode)
	trans(heard)
}

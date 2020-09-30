package main

import (
	"encoding/json"
	"fmt"
)

/*
	Golang JSON序列化是指把结构体数据转化成JSON格式的字符串
	Golang JSON的反序列化是指把JSON数据转化成Golang中的结构体对象

	json.marshal	---> 结构体转换成json字符串
	json.unmarshal	---> json字符串幻化成结构体
*/

//Student 学生
type Student struct {
	Id     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []Student
}

func main() {
	str := `{"Title":"001班","Students":[{"Id":1,"Gender":"男","Name":"stu_1"},{"Id":2,"Gender":"男","Name":"stu_2"},{"Id":3,"Gender":"男","Name":"stu_3"},{"Id":4,"Gender":"男","Name":"stu_4"},{"Id":5,"Gender":"男","Name":"stu_5"},{"Id":6,"Gender":"男","Name":"stu_6"},{"Id":7,"Gender":"男","Name":"stu_7"},{"Id":8,"Gender":"男","Name":"stu_8"},{"Id":9,"Gender":"男","Name":"stu_9"},{"Id":10,"Gender":"男","Name":"stu_10"}]}`
	c1 := &Class{} //这里的c1是指针变量
	err := json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("Json Unmarshal Failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
	fmt.Printf("%#v\n", c1.Title)
}

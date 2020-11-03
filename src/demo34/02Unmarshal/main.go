package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {

	//json字符串
	var str = `{"ID":12,"Gender":"男","Name":"李四","Sno":"s0001"}`

	var s1 Student
	err := json.Unmarshal([]byte(str), &s1) //需要传入指针变量"&s1"

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s1)
	fmt.Printf("%T\n", s1)
	fmt.Println(s1.Name)
}

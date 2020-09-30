package main

import (
	"encoding/json"
	"fmt"
)

//结构体标签(在把结构体转换成json数据串的时候,里面的字段可以转换成周字母为小写的)
type Student struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {

	var s1 = Student{
		Id:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v\n", s1)

	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr)
}

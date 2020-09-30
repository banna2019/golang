package main

import (
	"encoding/json"
	"fmt"
)

//如果要把结构体中的字段转换成json字段的话,结构体中的字段首字母必须大写
type Student struct {
	ID     int
	Gender string
	Name   string //私有属性不能被json包访问
	// name   string //私有属性不能被json包访问
	Sno string
}

func main() {
	//结构体实例化
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}

	fmt.Printf("%#v\n", s1)

	jsonByte, _ := json.Marshal(s1) //这里返回是一个Byte类型的切片

	fmt.Println(jsonByte) //这里返回是一个Byte类型的切片,需要把Byte切换转换成string类型
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr)
}

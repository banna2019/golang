package main

import "fmt"

func main() {
	var userinfo = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	// fmt.Println(userinfo["username"])

	//for range循环遍历map类型的数据
	for k, v := range userinfo {
		// fmt.Println(k, v)
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}

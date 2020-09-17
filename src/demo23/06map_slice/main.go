package main

import "fmt"

func main() {
	//如果想在切片里面放一系列用户的信息,这时候就可以定义一个元素为map类型的切片
	//map[]类型的值也可以是切片

	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["hobby"] = "睡觉"

	var userinfo = make(map[string][]string)
	userinfo["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}

	userinfo["work"] = []string{
		"php",
		"golang",
		"前端",
	}

	// fmt.Println(userinfo)
	for _, v := range userinfo {
		// fmt.Println(v)
		for key, value := range v {
			fmt.Printf("key:%v - value:%v\n", key, value)
		}
	}
}

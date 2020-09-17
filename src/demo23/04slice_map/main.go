package main

import "fmt"

func main() {
	//如果想在切片里面放一系列用户的信息,这时候就可以定义一个元素为map类型的切片
	//切片的元素可以为map[]类型

	var userinfo = make([]map[string]string, 3, 3)
	// fmt.Println(userinfo[0])        //map[] map不初始化的默认值nil
	// fmt.Println(userinfo[0] == nil) //true

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"

	}
	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "22"
		userinfo[1]["height"] = "170cm"

	}
	fmt.Println(userinfo) //如果设定了切片的长度和容量,数据写入的时候没有满足;切片的长度和容量会被不空map[]
}

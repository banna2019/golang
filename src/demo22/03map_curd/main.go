package main

import "fmt"

func main() {
	//map类型数据的curd

	//1.创建 修改map类型的数据
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// fmt.Println(userinfo)

	//2.创建 修改map类型的数据
	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// }
	// userinfo["username"] = "李四"
	// fmt.Println(userinfo)

	//3.获取查找map类型的数据
	/*
		var userinfo = map[string]string{
			"username": "张三",
			"age":      "20",
		}

		fmt.Println(userinfo["username"]) //获取

		//查找map类型中的数据
		v, ok := userinfo["age"]
		fmt.Println(v, ok) //20 true

		// v, ok := userinfo["xxxx"]
		// fmt.Println(v, ok) //空 和 false
	*/

	//4.删除map数据里面的可以及其对应的值
	/*
		delete(map 对象, key)
		​	• map 对象:	表示要删除键值对的map对象
		​	• key:	表示要删除的键值对的键
	*/

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// 	"height":   "180cm",
	// }
	// fmt.Println(userinfo)

	// delete(userinfo, "sex")
	// fmt.Println(userinfo)

	//5.map类型中新增数据(默认是覆盖的,如果不想覆盖,可以先判断一下键是否存在)
	var userinfo = map[string]string{
		"username": "张三",
		"age":      "20",
	}
	userinfo["username"] = "李四"
	fmt.Println(userinfo)

	// if _, ok := userinfo["height"]; !ok {
	// 	userinfo["height"] = "180cm"
	// }
	// fmt.Println(userinfo)
}

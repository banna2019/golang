package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

//Golang中空接口和类型断言使用细节
func main() {

	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["username"])
	fmt.Println(userinfo["age"])
	// fmt.Println(userinfo["hobby"][1]) //invalid operation: userinfo["hobby"][1] (type interface {} does not support indexing)

	var address = Address{
		Name:  "李四",
		Phone: 15543543983,
	}
	fmt.Println(address.Name)

	userinfo["address"] = address
	fmt.Println(userinfo["address"]) //{李四 15543543983}

	// var name = userinfo["address"].Name //userinfo["address"].Name undefined (type interface {} is interface with no methods)
	// fmt.Println(name)

	hobby1, _ := userinfo["hobby"].([]string) //空接口与类型断言的结合使用
	fmt.Println(hobby1[1])                    //吃饭

	address1, _ := userinfo["address"].(Address) //这里断言的是Address结构体
	fmt.Println(address1.Name, address1.Phone)
}

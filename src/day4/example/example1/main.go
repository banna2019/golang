package main

import (
	"errors"
	"time"
)

/*值类型*/
// func main() {
// 	var i int
// 	fmt.Println(i)

// 	j := new(int) //返回一个地址
// 	*j = 100
// 	fmt.Println(*j)

// 	//"*"获取内存地址中的值
// 	//"&"引用内存地址
// }

/*初始化一个error实例*/
func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	*/

	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	for {
		test()
		time.Sleep(time.Second)
	}
}

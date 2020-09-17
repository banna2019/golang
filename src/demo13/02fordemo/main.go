package main

import "fmt"

func main() {
	/*
		for 初始语句;条件表达式;结束语句{
			循环体语句
		}
	*/

	//1.打印1~10的所有珊瑚橘
	/*
		1.  i := 1
		2. 	i <= 10
		3.	执行花括号里面的语句
		4.	i++
		5.	i <= 10
		6.	执行花括号里面的语句
		7.	i++
		8.	i <= 10
	*/
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 1
	// for ; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	/*
		写for循环的时候要注意死循环
		for循环的初始语句和结束语句都可以省略,例如:
	*/

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	/*
		for {
			循环体语句
		}
		注意: Go 语言中是没有while语句的,可以通过for代替
	*/

	i := 1
	for {
		if i <= 10 {
			fmt.Println(i)
		} else {
			break //跳出循环
		}
		i++
	}

}

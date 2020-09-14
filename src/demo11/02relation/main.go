//关系运算
package main

import "fmt"

func main() {
	/*
	 ==		检查两个值是否相等,如果相等返回true,否则返回false
	 !=		检查两个值是否不相等,如果不相等返回true,否则返回false
	 >		检查左边值是否大于右边值,如果是返回true,否则返回false
	 >=		检查左边值是否大于等于右边值,如果是返回true,否则返回false
	 <		检查左边值是否小于右边值,如果是返回true,否则返回false
	 <=		检查左边值是否小于等于右边值,如果是返回true,否则返回false
	*/

	//演示关系运算的使用
	// var a1 = 9
	// var a2 = 8

	// fmt.Println(a1 == a2) //false
	// fmt.Println(a1 != a2) //true
	// fmt.Println(a1 > a2)  //true
	// fmt.Println(a1 >= a2) //true
	// fmt.Println(a1 < a2)  //false
	// fmt.Println(a1 <= a2) //false

	// var a1 = 9
	// var a2 = 8

	// flag := a1 > a2
	// if flag {
	// 	fmt.Println("a1 > a2")
	// }

	var a1 = 9
	var a2 = 8

	if a1 == a2 {
		fmt.Println("a1=a2")
	} else {
		fmt.Println("a1!=a2")
	}
}

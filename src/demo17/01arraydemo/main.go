package main

import "fmt"

func main() {

	//1.数组的长度类型的一部分
	// var arr1 [3]int
	// var arr2 [4]int
	// var strArr [3]string

	// fmt.Printf("arr1:%T arr2:%T strArr:%T", arr1, arr2, strArr)

	//2.数组的初始化 第一中方法
	// var arr1 [3]int
	// // fmt.Println(arr1)

	// arr1[0] = 23
	// arr1[1] = 10
	// arr1[2] = 24
	// fmt.Println(arr1)

	// var strArr [3]string

	// strArr[0] = "php"
	// strArr[1] = "java"
	// strArr[2] = "golang"
	// fmt.Println(strArr)

	//3.数组的初始化 第二种方法
	// var arr1 = [3]int{23, 34, 5}
	// fmt.Println(arr1)

	// var arr2 = [3]string{"java", "php", "golang"}
	// fmt.Println(arr2)

	// arr2 := [3]string{"java", "php", "golang"}
	// fmt.Println(arr2)

	//3.数组的初始化 第三种方法	一般情况下可以让编译器根据初始值的个数自行推断数组的长度
	// var arr1 = [...]int{12, 3344, 5343, 23, 434, 63}
	// fmt.Println(arr1)
	// fmt.Println(len(arr1))

	//注意数组的长度
	// arr2 := [...]string{"java", "php", "golang", "js"}
	// fmt.Println(arr2)

	//改变数组中元素的值
	// arr2 := [...]string{"java", "php", "golang", "js"}
	// arr2[0] = "phper"
	// fmt.Println(arr2)

	// 4.数组的初始化 第四种方法
	// arr := [...]int{0: 1, 1: 23, 2: 2, 3: 12, 4: 5, 5: 50}
	// fmt.Println(arr)
	// fmt.Println(len(arr))

	//5.数组的循环遍历for for range
	// var arr1 = [3]int{23, 34, 5}
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Println(arr1[i])
	// }

	// arr2 := [...]string{"java", "php", "golang", "js"}
	// for i := 0; i < len(arr2); i++ {
	// 	fmt.Println(arr2[i])
	// }

	arr2 := [...]string{"java", "php", "golang", "js"}
	for k, v := range arr2 {
		fmt.Printf("key:%v  value:%v\n", k, v)
	}
}

package main

import "fmt"

func main() {

	// 1.回顾数组的声明及初始化
	// var arr1 [3]int
	// arr1[0] = 10
	// arr1[1] = 20
	// arr1[2] = 30
	// fmt.Println(arr1)

	// var arr2 = [3]string{"php", "java", "golang"}
	// fmt.Println(arr2)

	// var arr3 = [...]string{"php", "java", "golang"}
	// fmt.Println(arr3)

	// var arr4 = [...]int{1: 2, 2: 3}
	// fmt.Println(arr4)

	//2.切片的声明初始化
	// var arr1 []int //不加长度就是声明一个切片
	// fmt.Printf("%v - %T - 长度:%v\n", arr1, arr1, len(arr1))

	// var arr2 = []int{1, 2, 3, 4, 45}
	// fmt.Printf("%v - %T - 长度:%v\n", arr2, arr2, len(arr2))

	// var arr3 = []int{1: 2, 2: 4, 3: 6}
	// fmt.Printf("%v - %T - 长度:%v", arr3, arr3, len(arr3))

	var arr1 []int
	var arr2 = []int{1, 2, 3, 4, 45}
	fmt.Println(arr1)
	fmt.Println(arr1 == nil) //true golang中声明切片后,切片的默认值是nil
	fmt.Println(arr2 == nil)
}

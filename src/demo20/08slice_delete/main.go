package main

import "fmt"

func main() {
	//Go语言中并没有删除切片元素的专用方法,可以使用切片本身的特殊性来删除元素
	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}

	// //要删除索引为2的元素 注意: append合并切片的时候最后一个元素要加...
	// a = append(a[:2], a[3:]...)
	// fmt.Println(a)

	sliceB := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//要删除35，36
	sliceB = append(sliceB[:5], sliceB[7:]...)
	fmt.Println(sliceB)

}

package main

import "fmt"

/*
	Go 语言的内建函数append()可以为切片动态添加元素,每个切片会指向一个底层数组,这个数组的容量够用就添加新增元素
*/

func main() {
	//1.append方法的使用
	//golang中给切片扩容的话用到append()方法
	// var sliceC []int
	// sliceC = append(sliceC, 12, 23, 35, 465)

	// fmt.Printf("%v -- %v--%v\n", sliceC, len(sliceC), cap(sliceC))
	// fmt.Println(sliceC)

	//2.append方法还可以合并切片
	// sliceA := []string{"php", "java"}
	// sliceB := []string{"nodejs", "go"}

	// sliceA = append(sliceA, sliceB...)
	// fmt.Println(sliceA)

	//3.切片的扩容策略
	var sliceA []int
	for i := 1; i < 10; i++ {
		sliceA = append(sliceA, i)
		fmt.Printf("%v 长度:%d 容量:%d\n", sliceA, len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)
}

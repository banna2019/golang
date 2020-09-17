package main

import "fmt"

func main() {
	/*
		什么叫做冒泡排序
			概念: 从头到尾,比较相邻的两个元素的大小,如果符合交换条件,交换两个元素的位置.
			特点: 每一轮比较中,都会选出一个最大的数,放在正确的位置
	*/

	//冒泡排序从小到大
	// var numSlice = []int{9, 6, 5, 4, 8}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := 0; j < len(numSlice)-1-i; j++ {
	// 		if numSlice[j] > numSlice[j+1] {
	// 			temp := numSlice[j]
	// 			numSlice[j] = numSlice[j+1]
	// 			numSlice[j+1] = temp
	// 		}
	// 	}
	// }
	// fmt.Println(numSlice)

	//冒泡排序从大到小
	var numSlice = []int{9, 6, 5, 4, 8}
	for i := 0; i < len(numSlice); i++ {
		for j := 0; j < len(numSlice)-1-i; j++ {
			if numSlice[j] < numSlice[j+1] {
				temp := numSlice[j]
				numSlice[j] = numSlice[j+1]
				numSlice[j+1] = temp
			}
		}
	}
	fmt.Println(numSlice)
}

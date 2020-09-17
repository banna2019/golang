package main

import "fmt"

func main() {
	//回顾练习: 从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2).

	// var arr = [...]int{1, 3, 5, 7, 8}
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[i]+arr[j] == 8 {
	// 			fmt.Printf("(%v,%v)\n", i, j)
	// 		}
	// 	}
	// }

	/*
		选择排序: 进行从小到大排序
		概念: 通过比较 首先选出最小的数放在第一位置上,然后在其余的数中选出次小数放在第二位置上,以此类推,直到所有的数成为有序序列
		[9,8,7,6,5,4]
	*/

	//从小到大排序
	// var numSlice = []int{9, 8, 7, 6, 5, 4, 12, 54, 23, 76}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := i + 1; j < len(numSlice); j++ {
	// 		if numSlice[i] > numSlice[j] {
	// 			temp := numSlice[i]
	// 			numSlice[i] = numSlice[j]
	// 			numSlice[j] = temp
	// 		}
	// 	}
	// }
	// fmt.Println(numSlice)

	//从到小的排序
	var numSlice = []int{9, 8, 7, 6, 5, 4, 12, 54, 23, 76}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] < numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	fmt.Println(numSlice)
}

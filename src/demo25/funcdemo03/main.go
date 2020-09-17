package main

import "fmt"

//int类型的升序排序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

//int类型的降序排序
func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

//int类型的降序排序
func sortIntDesc1(slice []int) { //没有返回值,直接进行调用
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
}

func main() {
	//案例1: 把前面讲的选择排序封装成方法,实现整型切片的升序降序排序排列
	// float和string类型的排序写法示例

	// var sliceA = []int{12, 34, 37, 35, 556, 36, 2}

	// arr := sortIntAsc(sliceA)
	// fmt.Println(arr)

	// var sliceB = []int{1, 34, 4, 35, 6, 36, 7}
	// fmt.Println(sortIntAsc(sliceB))

	// var sliceC = []int{12, 34, 37, 35, 556, 36, 2}
	// fmt.Println(sortIntDesc(sliceC))

	var sliceC = []int{12, 34, 37, 35, 556, 36, 2}
	sortIntDesc1(sliceC) //调用完"sortIntDesc1()"方法之后,这里的sliceC的数据已经发生变化了
	fmt.Println(sliceC)

	/*
		案例2:
			m1 := map[string]string{
				"username":"zhangsan",
				"age":"20",
				"sex":"男",
				"height":"180",
			}
			输出: age => 20height=>180sex=>男username=>zhangsan
			上面有一个map对象m1,封装一个方法,要求按照key升序排序,最后输出一个key=>value
	*/
}

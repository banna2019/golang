package main

func main() {
	// //修改切片中的值
	// var selicA = make([]int, 4, 8)
	// // fmt.Println(selicA)
	// selicA[0] = 10
	// selicA[1] = 12
	// selicA[2] = 40
	// selicA[3] = 30
	// fmt.Println(selicA)

	// sliceB := []string{"php", "java", "go"}

	// sliceB[2] = "golang"
	// fmt.Println(sliceB)

	//golang中没有办法通过下标的方式给切片扩容
	// var sliceC []int
	// fmt.Printf("%v -- %v--%v", sliceC, len(sliceC), cap(sliceC))
	// sliceC[0] = 1
	// fmt.Println(sliceC)

	//golang中给切片扩容的话用到append()方法
}

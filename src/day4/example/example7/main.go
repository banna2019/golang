package main

import "fmt"

/*数组小标*/
/*
func testSlice() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	slice = arr[2:5]
	slice = arr[:3]              //包含0到线标为3的所有的元素
	slice = arr[:]               //整个数组的元素
	slice = arr[1:]              //包含下标1,到整个数组的所有元素
	slice = slice[1:]            //把数组的第一个元素给去掉
	slice = slice[:len(slice)-1] //把数组的最后一个元素给去掉
	fmt.Println(slice)
	fmt.Println(len(slice)) //长度
	fmt.Println(cap(slice)) //容量

	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func main() {
	testSlice()
}*/

/*结构体*/
type slice struct { //这里自定义slice也是引用类型
	ptr *[100]int //这里是地址的指针
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice() {
	var s1 slice
	s1 = make1(s1, 10)

	s1.ptr[0] = 100
	s1.ptr[1] = 1000

	fmt.Println(s1.ptr)
}

func modify1(a []int) {
	a[1] = 1000
}

func testSlice1() {
	var b []int = []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func testSlice2() {
	var a = [10]int{1, 2, 3, 4, 5}

	b := a[1:5] //切片就是一个地址,切片传递的就是一个指向的地址
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])
}

func main() {
	// testSlice()
	testSlice1()
	testSlice2()
}

package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]

	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)

	//数据在达到上限之后,程序会开辟一块新的内存空间把之前的数组中的元素拷贝到新的内存空间中;再把新增加的元素也添加到新的内存空间中

	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
}

func testCopy() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 10) //这里b的容量为10

	copy(b, a)

	fmt.Println(b)
}

func testCopy1() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 1) //这里b的容量为1

	copy(b, a)

	fmt.Println(b)
}

/*string型是不可变的*/
func testString() {
	s := "hello world"
	s1 := s[0:5]
	s2 := s[6:]

	fmt.Println(s1)
	fmt.Println(s2)
}

func testModifyString() {
	s := "hello world"
	s1 := []rune(s)

	s1[1] = '0'
	str := string(s1)
	fmt.Println(str)
}

func main() {
	// testSlice()
	// testCopy()
	// testCopy1()
	// testString()

	testModifyString()
}

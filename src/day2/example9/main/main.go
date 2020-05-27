package main

import (
	"fmt"
)

/*第一种写法*/
// func swap(a *int, b *int) { //"*int"传指针
// 	tmp := *a //"*a"就是取的指针所指的内存区域中所存储的值
// 	*a = *b
// 	*b = tmp
// 	return
// }

// func main() {
// 	first := 100
// 	second := 200
// 	swap(&first, &second) //因为传入的是内存地址,所以这里需要使用"&"
// 	fmt.Println("first=", first)
// 	fmt.Println("second=", second)
// }

/*第二种写法*/
func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	first := 100
	second := 200
	// first, second = swap(first, second)
	first, second = second, first
	fmt.Println("first=", first)
	fmt.Println("second=", second)

	test2()
}

func test2() {
	var a int8 = 100
	// var b int16 = a //这里a和b虽然都是整数,但是它们都是不同类型;不同类型之间是不能直接进行转换的
	var b int16 = int16(a) //这里对a的类型进行转换

	fmt.Printf("a=%d b=%d\n", a, b)
}

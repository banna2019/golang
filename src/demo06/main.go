package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1.定义int类
	// var num int = 10

	// fmt.Printf("num=%v 类型：%T", num, num)

	// var num int8 = 98
	// fmt.Printf("num=%v 类型:%T", num, num)

	//2.int8的范围演示

	//(-127 到 127) 错误
	// var num int8 = 130
	// fmt.Printf("num=%v 类型:%T", num, num)

	// var num int16 = 130
	// fmt.Printf("num=%v 类型:%T", num, num)

	/*
		// unsafe.Sizeof查看不同长度的整型 在内存里面的存储空间
		var a int8 = 15
		// var a int16 = 15
		fmt.Printf("num=%v 类型:%T\n", a, a)
		fmt.Println(unsafe.Sizeof(a)) //存储空间 1 字节
	*/

	//3.uint8的范围(0-255)
	// var n1 uint8 = 200
	// // var n1 uint8 = -2 	//uint8 是无符号int类型
	// fmt.Printf("n1=%v 类型:%T", n1, n1)

	//4.int8 int16 ... 占用的存储空间大小
	// var a int8 = 15
	// var a int16 = 15
	// fmt.Printf("num=%v 类型:%T\n", a, a)
	// fmt.Println(unsafe.Sizeof(a)) //存储空间 1 字节

	//5.int不同长度直接的转换(int类型转换)
	// var a1 int32 = 10
	// var a2 int64 = 21

	// fmt.Println(int64(a1) + a2)
	// fmt.Println(a1 + int32(a2))

	//高位向低位转换时候需要注意
	var n1 int16 = 130
	fmt.Println(int8(n1)) //-126 有问题了

	var n2 int16 = 110
	fmt.Println(int8(n2))
	fmt.Println(int64(n2))

	//6.int数字字面语法 %d 表示10进制输出 %b表示二进制输出 %o 表示八进制输出 %x 表示十六进制输出

	num := 9

	fmt.Printf("num=%v 类型: %T\n", num, num)
	fmt.Println(unsafe.Sizeof(num)) //表示64位的计算机 int就是int64 占用8个字节

	fmt.Printf("num=%v\n", num) //%v 原样输出
	fmt.Printf("num=%d\n", num) //%d 表示10进制输出
	fmt.Printf("num=%b\n", num) //%b表示二进制输出
	fmt.Printf("num=%o\n", num) //%o 表示八进制输出
	fmt.Printf("num=%x\n", num) //%x 表示十六进制输出

}

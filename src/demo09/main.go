package main

import "fmt"

func main() {

	//1.golang中定义字符 字符属于int类型
	//单引号''字符,双引号""字符串
	//gong中"字符"是属于int类型,组成字符串的每个元素叫做"字符"

	// var a = 'a'
	// fmt.Printf("值: %v  类型: %T", a, a) //当直接输出byte(字符)的时候输出的是这个字符对应的码值

	//2.原样输出字符
	// var a = 'a'
	// fmt.Printf("值: %c  类型: %T", a, a)

	/*
		Go 语言的字符有以下两种:

		1.uint8 类型,或者叫byte型,代表了ASCII 码的一个字符.(byte类型主要处理ASCII的字符)

		2.rune 类型,代表一个UTF-8字符.(rune类型主要处理UTF8类型的字符,rune也兼容byte类型的字符串)

		​	当需要处理中文、日文或者其他复合字符时,则需要用到rune类型.rune类型实际是一个int32.

		​	Go使用了特殊的rune类型来处理Unicode,让基于Unicode的文本处理更为方便,也可以使用byte型进行默认字符串处理,性能和扩展性都有照顾.
	*/

	//3.定义一个字符串输出字符串里面的字符
	// var str = "this"
	// fmt.Printf("值: %v 原样输出: %c 类型: %T", str[2], str[2], str[2])

	//4.一个汉字占用3个字节(utf-8),一个字母占用一个字节

	//unsafe.Sizeof() 没法查看string类型数据占用的存储空间

	// var str = "this" //占用4个字节
	// // fmt.Println(unsafe.Sizeof(str))
	// fmt.Println(len(str))

	// var str = "你好go"
	// fmt.Println(len(str))

	//5.定义一个字符, 字符的值是汉字
	//golang中汉字使用的是utf8编码 编码后的值就是int类型
	// var a = '国'
	// fmt.Printf("值: %v  原样输出: %c  类型: %T", a, a, a) //值: 22269  原样输出: 国  类型: int32

	//6.通过循环输出字符串连的字符
	// str := "你好 golang"
	// for i := 0; i < len(str); i++ {		//for循环相当于byte
	// 	fmt.Printf("%v(%c)\n", str[i], str[i])
	// }

	// str := "你好 golang"
	// for _, v := range str { //range相当于rune
	// 	fmt.Printf("%v(%c)\n", v, v)
	// }

	//7.修改字符串
	// 要修改字符串,需要先将其转换成[]rune 或[]byte,完成后再转换为string.无论哪种转换,都会重新分配内存,并复制字节数组.

	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	s1 := "你好golang"
	runeStr := []rune(s1)
	runeStr[0] = '大'
	fmt.Println(string(runeStr))

}

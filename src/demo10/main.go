package main

import (
	"fmt"
)

func main() {
	//1.整型和整型之间转换
	// var a int8 = 20
	// var b int16 = 40

	// fmt.Println(int16(a) + b) //低位转高位

	//2.浮点型和浮点型之间的转换
	// var a float32 = 20
	// var b float64 = 40
	// fmt.Println(float64(a) + b)

	//3.整型和浮点型之间的转换
	// var a float32 = 20.23
	// var b int = 40
	// fmt.Println(a + float32(b))

	//注意: 转换的时候建议从低位转换成高位,高位转换成低位的时候如果转换不成功就会溢出,和想的结果并不一样

	//4.通过fmt.Sprintf()把其他类型转换成string
	//注意: Sprintf 使用中需要注意转换的格式 int为%d float为%f bool为%t byte为%c
	// var i int = 20
	// var f float64 = 12.456
	// var t bool = true
	// var b byte = 'a'

	// str1 := fmt.Sprintf("%d", i)
	// fmt.Printf("%v--类型:%T\n", str1, str1)

	// str2 := fmt.Sprintf("%.2f", f)
	// fmt.Printf("%v--类型%T\n", str2, str2)

	// str3 := fmt.Sprintf("%t", t)
	// fmt.Printf("%v--类型%T\n", str3, str3)

	// str4 := fmt.Sprintf("%c", b)
	// fmt.Printf("%v--类型%T", str4, str4)

	//5.使用strconv包里面的几种转换的方法,把其他类型转换成string类型
	/*
		FormatInt
		参数1: int64 的数值
		参数2: 传值int类型的进制
	*/
	// var i int = 20
	// str1 := strconv.FormatInt(int64(i), 10)
	// fmt.Printf("%v--类型:%T", str1, str1)

	/*
		参数1: 要转换的值
		参数2: 格式化类型'f' (-ddd.ddd)、
			’b‘ (-ddddp±ddd,指数为二进制)、
			’e‘ (-d.dddde±dd,十进制指数)、
			’E‘ (-d.ddddE±dd,十进制指数)、
			’g‘ (指数很大时用'e'格式,否则'f'格式)、
			’G‘ (指数很大时用'E'格式,否则'f'格式)
		参数3: 保留的小数点-1(不对小数点格式化)
		参数4: 格式化的类型传入 64 32
	*/
	// var f1 float32 = 20.113123
	// str2 := strconv.FormatFloat(float64(f1), 'f', 4, 64)
	// fmt.Printf("%v--类型:%T", str2, str2)

	//6.ForatBool将bool转换为string
	// str3 := strconv.FormatBool(true) //没有任何意义
	// fmt.Printf("%v--类型:%T", str3, str3)

	// a := 'a' //意义不大
	// str4 := strconv.FormatUint(uint64(a), 10)
	// fmt.Printf("%v--类型:%T", str4, str4)

	//7.String类型转换成int整型
	/*
		ParseInt
			参数1: string数据
			参数2: 进制
			参数3: 位数 32 64 16
	*/
	// str := "1234"
	// fmt.Printf("%v--%T\n", str, str)
	// num, _ := strconv.ParseInt(str, 10, 64)
	// fmt.Printf("%v--类型:%T", num, num)

	/*
		ParseInt
			参数1: string数据
			参数2: 位数 32 64 16
	*/
	// str := "1234"

	// num, _ := strconv.ParseFloat(str, 64)
	// fmt.Printf("%v--类型:%T", num, num)

	//不建议把string类型转换成bool类型
	// b, _ := strconv.ParseBool("true") //string转bool
	// fmt.Printf("%v--类型:%T", b, b)

	s := "hello 张三"
	for _, r := range s {
		fmt.Printf("%v(%c)\n", r, r)
	}

}

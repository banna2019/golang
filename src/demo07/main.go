package main

import "fmt"

func main() {
	/*
		//1.定义float类型

		// var a float32 = 3.12
		// fmt.Printf("值: %v--%f,类型:%T\n", a, a, a) //值: 3.12--3.120000,类型:float32
		// fmt.Println(unsafe.Sizeof(a))            //float32占用4个字节

		var b float64 = 3.12
		fmt.Printf("值: %v--%f,类型:%T\n", b, b, b) //值: 3.12--3.120000,类型:float64
		fmt.Println(unsafe.Sizeof(b))            //float32占用4个字节
	*/

	/*
		//2.%f 输出float类型  %.2f 输出数据的时候保留2位小数
		var c float64 = 3.1415925535
		fmt.Printf("%v--%f--%.2f--%.4f", c, c, c, c)
	*/

	/*
		//3. 64位系统中 浮点数默认是 float64

		// f1 := 3.1345656456
		// fmt.Printf("%v--%T", f1, f1)

		//4. Golang 科学计数法表示浮点类型

		// var f2 float32 = 3.14e2 //表示f2等于3.14*10的2次方
		// fmt.Printf("%f--%T\n", f2, f2)
		// fmt.Printf("%v--%T", f2, f2)

		// var f3 float32 = 3.14e-2 //表示f3等于3.14除以10的2次方
		// fmt.Printf("%f--%T\n", f3, f3)
		// fmt.Printf("%v--%T", f3, f3)

		//5. Golang中 float 精度丢失问题(解决办法引用第三方包)

		// var f4 float64 = 1129.6
		// fmt.Println(f4 * 100) //输出:112959.99999999999

		// m1 := 8.2
		// m2 := 3.8

		// fmt.Println(m1 - m2) //期望是 4.4,结果打印出了 4.3999999999999995

		//6.int类型转换成float类型

		// a := 10
		// b := float64(a)
		// fmt.Printf("a的类型是%T,b类型是%T", a, b)

		//7. float32类型转换成float64类型
		// var a1 float32 = 23.4
		// a2 := float64(a1)
		// fmt.Printf("a1的类型是%T,a2的类型是%T", a1, a2)

		//8.float类型转换成int类型(不建议这样操作)

		// var c1 float32 = 23.45
		// c2 := int(c1)
		// fmt.Printf("c2的值: %v  c2的类型:%T", c2, c2)
	*/

	/*
		Go语言中bool类型进行声明布尔类型数据,布尔类型数只有true(真)和false(假)两个值.

		注意:
		1.布尔类型变量的默认值为false.
		2.Go 语言中不允许将整型强制转换为布尔类型.
		3.布尔型无法参与数值运算,也无法与其他类型进行转换.

	*/

	// var flag bool = true
	// // var flag = true
	// fmt.Printf("%v--%T", flag, flag)

	//1.布尔类型变量默认值为false
	// var b bool
	// fmt.Printf("%v\n", b)

	//2.string类型变量的默认值为空
	// var s string
	// fmt.Printf("%v", s)

	//3.int类型变量的默认值为0
	// var i int
	// fmt.Printf("%v", i)

	//4.float类型的默认值为0
	// var f float32
	// fmt.Printf("%v", f)

	//5. Go语言中不允许将整型强制转换为布尔类型.
	// var a = 1
	// var a = true

	// if a { //错误写法
	// 	fmt.Printf("true")
	// }

	//6.布尔类型无法参与数值运算,也无法与其他类型进行转换
	// var s = "this is str"
	// if s { //错误写法
	// 	fmt.Printf("true")
	// }

	var f1 = false
	if f1 { //正确写法
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}

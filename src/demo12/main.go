package main

import "fmt"

/*
//练习题
func main() {

	//练习题1: 有两个变量,a和b,要求将其进行交换,最终打印结果

	// var a = 34
	// var b = 10

	// t := a
	// a = b
	// b = t
	// fmt.Printf("a=%v b=%v", a, b)

	//练习题2: 有两个变量,a和b,要求将其进行交换(不能使用中间变量),最终打印结果
	// var a = 34
	// var b = 10

	// a = a + b
	// b = a - b
	// a = a - b
	// fmt.Printf("a=%v b=%v", a, b)

	//练习题3: 假如还有100天放假,问: xx个星期零xx天
	// var days = 100
	// var week = days / 7
	// var day = days % 7
	// fmt.Printf("week=%v day=%v\n", week, day)
	// fmt.Printf("距离放假还有%v周,零几天%v", week, day)

	//练习题4: 定义一个变量保存华氏温度,华氏温度转换摄氏温度的公式为: C=(F-32)➗1.8,摄氏温度和华氏温度,请求出华氏温度对应的摄氏温度
	var F float32 = 100 //华氏温度
	C := (F - 32) / 1.8
	fmt.Printf("华氏温度对应的摄氏温度%.2f", C)
}
*/

//为运算(位运算符对整数在内存中的二进制位进行操作)
func main() {
	/*
	   &	两位为1才为1
	   |	两位有一个为1就为1
	   ^	相异或两位不一样则为1
	   <<	左移n位就是乘以2的n次方.  "a<<b"是把a的各二进为全部左移b位,高位丢弃,低位补0.
	   >>	右移n位就是除以2的n次方.
	*/

	var a = 5
	var b = 2

	fmt.Println("a&b", a&b)   //000 	值0
	fmt.Println("a|b", a|b)   //111 	值7
	fmt.Println("a^b", a^b)   //111 	值7
	fmt.Println("a<<b", a<<b) //10100 	值20	5*2的2次方
	fmt.Println("a>>b", a>>b) // 	值1	5/2的2次方

}

package main

import (
	"fmt"
	"strings"
)

func main() {

	//1.定义string类型

	// var str1 string = "你好golang"
	// var str2 = "你好 go"
	// str3 := "你好golang"

	// fmt.Printf("%v--%T \n", str1, str1)
	// fmt.Printf("%v--%T \n", str2, str2)
	// fmt.Printf("%v--%T \n", str3, str3)

	//2.转义符

	// str1 := "this \nis str" //换行
	// fmt.Println(str1)

	// str2 := "/Users/banna/Documents/Github/golang" //路径输出转义,window需要在反斜杠处,再加一个反斜杠
	// fmt.Println(str2)

	// str2 := "C:Go\"bin"
	// fmt.Println(str2)

	//3.多行字符串 ``(反引号) tab键上方

	// str1 := `this is str
	// this is str
	// this is str
	// this is str`

	// fmt.Println(str1)

	//4.len(str1) 求长度

	// var str1 = "你好"
	// fmt.Println(str1)
	// fmt.Println(len(str1))

	//5. + 或者 fmt.Sprintf拼接字符串

	// str1 := "你好"
	// str2 := "golang"
	// str3 := str1 + str2
	// fmt.Println(str3)

	// str1 := "你好"
	// str2 := "golang"
	// str3 := fmt.Sprintf("%v %v", str1, str2) 	//这里的Sprintf是直接返回内容的,需要被赋值给对象
	// fmt.Println(str3)

	// str1 := "反引号间换行将被作为字符串中的换行," +
	// 	"但是所有的转义字符均无效," +
	// 	"文本将会原样输出."

	// fmt.Println(str1)

	//6.string.Split分割字符串,strings需要引入strings包

	// var str = "123-456-789"
	// arr := strings.Split(str, "-")
	// fmt.Println(arr) //[123 456 789] 切片,在golang中切片和数组还有一些区别

	//7.strings.Join(a[]string,sep string) join操作 表示把切片链接成字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// str2 := strings.Join(arr, "*")
	// fmt.Println(str2)

	// arr := []string{"php", "java", "golang"} //切片
	// // fmt.Println(arr)
	// str3 := strings.Join(arr, "-")
	// // fmt.Println(str3)
	// fmt.Printf("%v - %T", str3, str3)

	//8.strings.contains判断是否包含(整个字符串的包含)
	// str1 := "this is str"
	// str2 := "thisxxx"
	// flag := strings.Contains(str1, str2)
	// fmt.Println(flag)

	//9.strings.HasPrefix,strings.HasSuffix 前缀/后缀判断
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.HasPrefix(str1, str2)
	// fmt.Println(flag)

	// str1 := "this is str"
	// // str2 := "this"
	// str2 := "str"
	// flag := strings.HasSuffix(str1, str2)
	// fmt.Println(flag)

	//10. stings.Index(),stings.LastIndex() 子串出现的位置,查找不到返回-1,查找到返回下标,下标是从0开始的

	// str1 := "this is str"
	// str2 := "is"
	// num := strings.Index(str1, str2) //Index的查找顺序是从前往后进行查找的
	// fmt.Println(num)

	// str1 := "this is str"
	// str2 := "is"
	// num := strings.LastIndex(str1, str2) //LastIndex的查找顺序是从后往前进行查找的
	// fmt.Println(num)	//返回5

	// str1 := "this is str"
	// str2 := "s"
	// num := strings.LastIndex(str1, str2) //LastIndex的查找顺序是从后往前进行查找的
	// fmt.Println(num)	//返回8

	str1 := "this is str"
	str2 := "xxx"
	num := strings.LastIndex(str1, str2)
	fmt.Println(num) //返回-1
}

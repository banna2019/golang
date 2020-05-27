package main

import "fmt"

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func reverse1(str string) string {
	var result []byte //数组是中括号在前面,后面是类型
	temp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, temp[length-i-1])
	}
	return string(result)

}

func main() {
	var srt1 = "Hello"
	var str2 = "World"
	var str3 = srt1 + " " + str2

	str4 := fmt.Sprintf("%s %s", srt1, str2)
	n := len(str4)

	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Printf("len(str4)=%d\n", n)

	substr := str4[0:5]
	fmt.Println(substr)

	substr = str4[6:] //因为是从7开始取小标到末尾,所以末尾可以不用写
	fmt.Println(substr)

	result := reverse(str4)
	fmt.Println(result)

	result = reverse1(str4)
	fmt.Println(result)

}

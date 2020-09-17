package main

import "fmt"

func main() {
	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	s2 := "你好golang"
	runeStr := []rune(s2)
	fmt.Println(runeStr)
	runeStr[0] = '大' //''分好为字节,""引号为字符串
	fmt.Println(string(runeStr))
}

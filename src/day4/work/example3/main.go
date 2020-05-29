package main

import "fmt"

func process(str string) bool {
	t := []rune(str) //"rune"无论是中文还是英文,都表示一个字符
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		// 调试代码段
		// fmt.Printf("%c---%c,%d---%d,%v",str[i],str[last],i,last,v)
		// fmt.Printf("%v-%v-%d\n",i,v,len(string(v)))
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

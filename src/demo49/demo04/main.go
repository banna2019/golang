package main

import (
	"fmt"
	"io/ioutil"
)

// ioutil按行读取示例(适用用于小文件内容读取,读取文件大小在100M~200M)
func main() {
	byteStr, err := ioutil.ReadFile("/Users/banna/Documents/Github/golang/src/demo48/08refect_struct/main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}

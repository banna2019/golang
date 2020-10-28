package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
二、读取文件（方法2）bufio 读取文件("bufio"是以流的方式读取的文件内容)

		1、只读方式打开文件 file,err := os.Open()

		2、创建reader对象  reader := bufio.NewReader(file)

		3、ReadString读取文件  line, err := reader.ReadString('\n')

		4、关闭文件流 defer file.Close()
*/

func main() {
	file, err := os.Open("/Users/banna/Documents/Github/golang/src/demo48/08refect_struct/main.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	//读取文件
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //一次读取一行
		if err == io.EOF {                  //使用"bufio.NewReader"来读取的话,注意标记完成以后;可能还会返回"str"
			fileStr += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
二、写入文件（方法2） bufio 写入文件

		1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

		2、创建writer对象  writer := bufio.NewWriter(file)

		3、将数据先写入缓存  writer.WriteString("你好golang\r\n")

		4、将缓存中的内容写入文件	writer.Flush()

		5、关闭文件流 file.Close()
*/
func main() {
	// file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("Open file failed,err:", err)
		return
	}
	write := bufio.NewWriter(file)
	// write.WriteString("你好golang") //表示将数据写入缓存中
	for i := 0; i <= 10; i++ {
		write.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n")
	}
	write.Flush() //将缓存中的数据写入文件

}

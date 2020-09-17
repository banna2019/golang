package main

import (
	"errors"
	"fmt"
)

/*
panic/recover

Go 语言中目前（Go1.12）是没有异常机制，但是使用 panic/recover 模式来处理错误。

panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效
*/

//模拟了一个文件读取的方法
func readFile(filename string) error {
	if filename == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("xxx.go")
	// err := readFile("main.go")
	if err != nil {
		panic(err)
	}
}

func main() {

	myFn()
	fmt.Println("继续执行...")
}

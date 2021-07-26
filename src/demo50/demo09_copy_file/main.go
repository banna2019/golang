package main

import (
	"fmt"
	"io"
	"os"
)

//文件式拷贝文件
//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(srcFileName string, dstFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	defer sFile.Close()
	dFile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer dFile.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	var tempSlice = make([]byte, 128)
	for {
		//读取数据
		n1, err := sFile.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//写入数据
		if _, err := dFile.Write(tempSlice[:n1]); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	//调用CopyFile 完成文件拷贝
	srcFile := "./main.go"
	dstFile := "./test.txt"
	err := CopyFile(srcFile, dstFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}

}
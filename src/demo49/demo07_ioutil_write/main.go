package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello golang"
	err := ioutil.WriteFile("./test1.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("Write file failed,err:", err)
		return
	}
}

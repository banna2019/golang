package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("./test1.txt", "./test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("重命名成功")
}

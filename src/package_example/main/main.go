package main

import (
	"fmt"
	"package_example/calc" //这里需要填写项目的相对路径才能引用自定义包,这里是需要填写"$GOPATH"的代码相对路径
)

func main() {
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 300)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}

package main

import "fmt"

func main() {
	//goto 语句通过标签进行代码间的无条件跳转.goto语句可以在快速跳出循环、避免重复退出上有一定帮助
	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto label3 //goto就直接跳过到label3之后执行了
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
label3:
	fmt.Println("ccc")
	fmt.Println("ddd")
}

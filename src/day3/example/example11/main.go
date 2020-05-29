package main

import "fmt"

func main() {
	var a int = 10

	switch a {
	case 0:
		// case 0，1,2,3,4,5:	//多条件
		fmt.Println("a  is equal 0")
		// fallthrough //执行完当前语句之后,还可以继续往下走
	case 10:
		fmt.Println("a is equal 10")
	default:
		fmt.Println("a is equal default")
	}
}

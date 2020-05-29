package main

import (
	"fmt"
)

/*直接传输值类型*/
// func modify(n int) {
// 	fmt.Println("in modify funciton,", &n)
// 	n = 100
// }

// func main() {
// 	var n int

// 	fmt.Println("out modify function,", &n)
// 	modify(n)
// 	fmt.Println(n)
// }

/*传入地址类型*/
func modify(n *int) {
	fmt.Println("in modify function,", n)
	*n = 110
}

func main() {
	var n int

	fmt.Println("out modify funciton,", &n)
	modify(&n)
	fmt.Println(n)
}

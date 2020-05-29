package main

import "fmt"

func modify(a int) {
	a = 100

}

func main() {
	var b int = 10
	a := &b
	// fmt.Println(*a)
	modify(*a)
	fmt.Println(*a)

}

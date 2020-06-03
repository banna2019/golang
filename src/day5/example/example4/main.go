package main

import "fmt"

func test(ap *int) {
	fmt.Println(*ap)
}

func main() {
	a := 10
	ap := &a

	test(ap)

}

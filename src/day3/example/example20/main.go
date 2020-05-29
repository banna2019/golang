package main

import "fmt"

/*方式一*/
// func test(a, b int) int {
// 	result := func(a1 int, b1 int) int {
// 		return a1 + b1
// 	}(a, b)
// 	return result
// }

/*方式二*/
// func test(a, b int) int {
// 	result := func(a1 int, b1 int) int {
// 		return a1 + b1
// 	}
// 	return result(a, b)
// }

// func main() {
// 	fmt.Println(test(100, 200))
// }

/*方式三*/
var (
	result = func(a1 int, b1 int) int {
		return a1 + b1
	}
)

func main() {

	fmt.Println(result(100, 200))
}

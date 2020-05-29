package main

import "fmt"

func main() {
	str := "hello.world,中国"

	for index, val := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	}
}

/*增强版*/
// func main() {
// 	str := "hello.world,中国"

// 	for index, val := range str {
// 		if index > 2 {
// 			continue
// 		}
// 		if index > 3 {
// 			break
// 		}
// 		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
// 	}
// }

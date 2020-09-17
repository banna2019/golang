package main

import "fmt"

func main() {
	//continue 语句可以结束当前循环,开始下一次的循环迭代过程,仅限在for循环中使用

	// for i := 1; i <= 10; i++ {
	// 	if i == 3 {
	// 		// break
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//在continue 语句后添加标签时,表示开始标签对应的循环
label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}
}

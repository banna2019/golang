package main

import (
	"fmt"
	"math"
)

/*优化前*/
// func isPrime(n int) bool {
// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

/*优化后*/
func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m) //"Scanf"从终端读取数据

	for i := n; i < m; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}

package main

import "fmt"

type calc func(int, int) int //表示定义一个calc的类型

type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test......")
}

func main() {
	// var c calc
	// var c1 calc
	// c = add
	// c1 = sub
	// fmt.Printf("c的类型: %T\n", c) //c的类型: main.calc
	// fmt.Println(c(10, 5))
	// fmt.Println(c1(10, 5))

	// f := sub                    //这里的f没有指定类型
	// fmt.Printf("f的类型: %T\n", f) //f的类型: func(int, int) int
	// fmt.Println(f(15, 5))

	var a int = 10
	var b myInt = 20
	fmt.Printf("a的类型: %T\n", a)
	fmt.Printf("b的类型: %T", b)
	fmt.Println(a + int(b)) //算数运算操作需要做类型转换,虽然b为myInt类型
}

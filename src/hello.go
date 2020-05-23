package main

// func main() {
// 	fmt.Println("hello world")
// }

func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func main() {
	//这里是单行注释
	// var c int
	// c = add(100, 200)
	// go test_goroute(300, 300)
	// fmt.Println("add(100,200)=", c)
	// fmt.Println("Hello World")

	//以下是多行注释
	/*
		for i := 0; i < 100; i++ {
			go test_print(i)
		}
		time.Sleep(10 * time.Second)
	*/

	test_pipe()

}

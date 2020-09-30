package main

func main() {
	//1.在计算机底层a这个变量其实对应了一个内存地址
	// var a int = 10
	// fmt.Printf("a的值: %v  a的类型%T a的地址%p", a, a, &a)

	//2.指针也是一个变量,但它是一种特殊的变量,它存储的数据不是一个普通的值,而是另一个变量的内存地址
	// var a int = 10
	// var p = &a //p指针类型	p的类型	*int(指针类型)

	// fmt.Printf("a的值: %v a的类型%T a的地址%p\n", a, a, &a)

	// fmt.Printf("p的值: %v p的类型%T\n", p, p)
	// fmt.Printf("p的值: %v p的类型%T", *p, p)

	//3.每一个变量都有自己的内存地址
	// var a = 10
	// var b = &a //b的值就是a的内存地址,b地址又是b自身的内存地址

	// fmt.Printf("a的值: %v a的类型%T  a的地址%p\n", a, a, &a)
	// fmt.Printf("b的值: %v b的类型%T b的地址%p\n", b, b, &b)

}

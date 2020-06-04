package main

import "fmt"

/*无冲突写法*/
/*
type Cart struct {
	name string
	age  int
}

type Train struct {
	Cart
	int
	start time.Time
}

func main() {
	var t Train

	//如果字段是结构体类型的,前面"Car"这个类型的字段可以不用写
	t.Cart.age = 200
	t.Cart.name = "001"

	//结构体字段的简写
	t.name = "Train"
	t.age = 100
	t.int = 200

	fmt.Println(t)
}*/

/*有冲突写法*/
type Car1 struct {
	name string
	age  int
}

type Car2 struct {
	name string
	age  int
}

type Train struct {
	Car1
	Car2
}

func main() {
	var t Train
	t.Car1.name = "Train"
	t.Car1.age = 100

	fmt.Println(t)
}

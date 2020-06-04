package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("Running")
}

type Bike struct {
	Car
	lunzi int
}

type Train struct {
	Car
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
	return str
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.lunzi = 2

	// fmt.Println(a)
	// a.Run()

	//普通继承
	// var b Train
	// b.weight = 100
	// b.name = "Train"
	// b.Run()

	//组合匿名字段的继承写法
	// var b Train
	// b.c.weight = 100
	// b.c.name = "Train"
	// b.c.Run()

	var b Train
	b.weight = 100
	b.name = "Train"
	b.Run()

	fmt.Printf("%s", &b)

}

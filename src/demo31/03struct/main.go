/*
Go语言中的基础数据类型可以表示一些事物的基本属性,
但是当想表达一个事物的全部或部分属性时,
这时候再用单一的基本数据类型明显就无法满足需求了,
Go语言提供了一种自定义数据类型,可以封装多个基本数据类型,
这种数据类型叫结构体,英文名称struct
*/

//注意: 结构体首字母大小写都可以,大写表示这个结构体是公有的,在其他包里面可以使用;小写表示这个结构体是私有的,只有这个包里才能使用

//注意: 在Golang中支持对结构体指针直接使用"."来访问结构体的成员.p2.name="张三"其实在底层是(*p2).name="张三"

package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {

	//实例化结构体2
	var p2 = new(Person)
	p2.Name = "李四"
	p2.Age = 20
	p2.Sex = "男"
	(*p2).Name = "王五"

	fmt.Printf("值: %#v  类型:%T\n", p2, p2)

	//实例化结构体3
	var p3 = &Person{} //"&"结构体引用
	p3.Name = "王五"
	p3.Age = 23
	p3.Sex = "男"

	fmt.Printf("值: %#v  类型:%T\n", p3, p3)

	//实例化结构体4
	var p4 = Person{
		Name: "哈哈",
		Age:  20,
		Sex:  "男",
	}

	fmt.Printf("值: %#v  类型:%T\n", p4, p4)

	//实例化结构体5
	var p5 = &Person{
		Name: "王麻子",
		Age:  20,
		Sex:  "男",
	}

	fmt.Printf("值: %#v  类型: %T\n", p5, p5) //这里输出的最后的类型为,引用类型

	//实例化结构体6
	var p6 = &Person{
		Name: "王麻子1",
	}
	fmt.Printf("值: %#v  类型: %T\n", p6, p6) //实例化的时候,struct中int类型不传递的话默认为"0",string不传递话默认为""

	//实例化结构体7
	var p7 = &Person{
		"张二",
		20,
		"男",
	}

	fmt.Printf("值: %#v  类型:%T\n", p7, p7)
}

package main

import "fmt"

func main() {

	//1.创建channel

	// ch := make(chan int, 3)

	//2.给管道里面存储数据
	// ch <- 10
	// ch <- 21
	// ch <- 32

	// //3.获取管道里面的内容
	// a := <-ch
	// fmt.Println(a)

	// <-ch //从管道里面取值,不赋值
	// c := <-ch
	// fmt.Println(c)

	// ch <- 56
	// ch <- 66
	// //4.打印管道的长度和容量
	// fmt.Printf("值: %v 容量: %v 长度: %v", ch, cap(ch), len(ch)) //值: 0xc0000ba000 容量: 3 长度: 1

	//5.管道的类型(引用数据类型)
	// ch1 := make(chan int, 4)

	// ch1 <- 34
	// ch1 <- 54
	// ch1 <- 64

	// ch2 := ch1
	// ch2 <- 69
	// <-ch1
	// <-ch1
	// <-ch1

	// d := <-ch1
	// fmt.Println(d)

	//6.管道阻塞
	// ch3 := make(chan int, 1) //fatal error: all goroutines are asleep - deadlock! 管道阻塞(管道容量满了之后,会阻塞)
	// ch3 <- 34
	// ch3 <- 64

	// ch4 := make(chan string, 2)
	// ch4 <- "数据1"
	// ch4 <- "数据2"

	// m1 := <-ch4
	// m2 := <-ch4
	// m3 := <-ch4
	// fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!	管道阻塞(管道里面没有数据了还在取数据也会阻塞)

	//正确写法
	ch5 := make(chan int, 1)
	ch5 <- 34
	<-ch5
	ch5 <- 67
	<-ch5
	ch5 <- 78
	m4 := <-ch5
	fmt.Println(m4)
}

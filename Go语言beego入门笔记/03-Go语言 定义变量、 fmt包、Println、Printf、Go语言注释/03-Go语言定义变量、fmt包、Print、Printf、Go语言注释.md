1、Go 语言定义变量...........................................................................................................1

2、fmt 包、Print、Println、Printf...................................................................................... 2

3、Go 语言中的注释...........................................................................................................3



### 1、Go 语言定义变量

关于变量:  程序运行过程中的数据都是保存在内存中,想要在代码中操作某个数据时就

需要去内存上找到这个变量,但是如果直接在代码中通过内存地址去操作变量的话,代

码的可读性会非常差而且还容易出错,所以利用变量将这个数据的内存地址保存起

来,以后直接通过这个变量就能找到内存上对应的数据了.



##### Golang中常见的变量定义方法如下:

###### 1、var定义变量

```go
var 变量名类型= 表达式

var name string = "zhangsan"
```



###### 2、类型推导方式定义变量

​	a在函数内部,可以使用更简略的 := 方式声明并初始化变量.

​	注意:  短变量只能用于声明局部变量,不能用于全局变量的声明

```go
变量名:= 表达式

n := 10
```



### 2、fmt 包、Print、Println、Printf

Go中要打印一个值需要引入fmt包

```go
import "fmt"
```

fmt包里面提供了一些常见的打印数据的方法,比如：Print 、Println、Printf在实际开发中Println、Printf用的非常多.



##### 1、Print和Println区别:

一次输入多个值的时候Println中间有空格Print没有

```go
fmt.Println("go", "python", "php", "javascript") // go python php javascript

fmt.Print("go", "python", "php", "javascript") // gopythonphpjavascript
```



Println会自动换行,Print不会

```go
package main
import "fmt"
func main() {
fmt.Println("hello")
fmt.Println("world")
// hello
// world
  fmt.Print("hello")
fmt.Print("world")
// helloworld
}
```



##### 2、Println和Printf区别

Printf是格式化输出,在很多场景下比Println更方便,举个例子:

```go
a := 10
b := 20
c := 30
fmt.Println("a=", a, ",b=", b, ",c=", c) //a= 10 ,b= 20 ,c= 30
fmt.Printf("a=%d,b=%d,c=%d", a, b, c) //a=10,b=20,c=30
```

%d是占位符,表示数字的十进制表示.Printf中的占位符与后面的数字变量一一对应.

更多的占位符参考: http://docscn.studygolang.com/pkg/fmt/



### 3、Go语言中的注释

win下面ctrl+/ 可以快速的注释一样,mac下面command+/也可以快速的注释一样

```go
/*
这是一个注释
*/
//这是一个注释
```


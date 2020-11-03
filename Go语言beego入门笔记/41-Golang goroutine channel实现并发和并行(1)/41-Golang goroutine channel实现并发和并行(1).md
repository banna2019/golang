一、为什么要使用goroutine..............................................................................................1

二、进程、线程以及并行、并发.......................................................................................2

三、Golang 中的协程（goroutine）以及主线程...............................................................4

四、Goroutine 的使用以及sync.WaitGroup......................................................................5

五、启动多个Goroutine.....................................................................................................7

六、设置Golang 并行运行的时候占用的cup 数量.............................................................8

七、Goroutine 统计素数....................................................................................................9

八、Channel 管道..............................................................................................................11

​	1、channel 类型.............................................................................................................11

​	2、创建channel..............................................................................................................11

​	3、channel 操作.............................................................................................................12

​	4、管道阻塞....................................................................................................................13

​	5、for range 从管道循环取值..........................................................................................14

九、Goroutine 结合Channel 管道........................................................................................15

十、单向管道...................................................................................................................... 20

十一、select 多路复用.........................................................................................................20

十二、Golang 并发安全和锁................................................................................................23

十三、Goroutine Recover 解决协程中出现的Panic............................................................28



### 一、为什么要使用goroutine

需求:  要统计1-10000000的数字中那些是素数,并打印这些素数？

素数:  就是除了1和它本身不能被其他数整除的数

实现方法:

1、传统方法,通过一个for循环判断各个数是不是素数

2、使用并发或者并行的方式,将统计素数的任务分配给多个goroutine去完成,这个时候就用到了goroutine

3、goroutine结合channel



### 二、进程、线程以及并行、并发

##### 1、关于进程和线程

进程(Process)就是程序在操作系统中的一次执行过程,是系统进行资源分配和调度的基

本单位,进程是一个动态概念,是程序在执行过程中分配和管理资源的基本单位,每一个进

程都有一个自己的地址空间.一个进程至少有5 种基本状态,它们是:  初始态,执行态,等待状态,就绪状态,终止状态.

通俗的讲进程就是一个正在执行的程序.

线程是进程的一个执行实例,是程序执行的最小单元,它是比进程更小的能独立运行的基本单位

一个进程可以创建多个线程,同一个进程中的多个线程可以并发执行,一个程序要运行的话至少有一个进程.



![image-20201002232729696](/Users/banna/Library/Application Support/typora-user-images/image-20201002232729696.png)



![image-20201002232751477](/Users/banna/Library/Application Support/typora-user-images/image-20201002232751477.png)



##### 2、关于并行和并发

并发:  多个线程同时竞争一个位置,竞争到的才可以执行,每一个时间段只有一个线程在执行.

并行:  多个线程可以同时执行,每一个时间段,可以有多个线程同时执行.

通俗的讲多线程程序在单核CPU上面运行就是并发,多线程程序在多核CUP上运行就是并

行,如果线程数大于CPU核数,则多线程程序在多个CPU上面运行既有并行又有并发



![image-20201002232851254](/Users/banna/Library/Application Support/typora-user-images/image-20201002232851254.png)



![image-20201002232908247](/Users/banna/Library/Application Support/typora-user-images/image-20201002232908247.png)



### 三、Golang中的协程(goroutine)以及主线程

golang中的主线程:  (可以理解为线程/也可以理解为进程),在一个Golang程序的主线程上可以起多个协程.Golang中多协程可以实现并行或者并发.

协程:  可以理解为用户级线程,这是对内核透明的,也就是系统并不知道有协程的存在,是

完全由用户自己的程序进行调度的.Golang的一大特色就是从语言层面原生支持协程,在

函数或者方法前面加go关键字就可创建一个协程.可以说Golang中的协程就是

goroutine .

![image-20201002233038391](/Users/banna/Library/Application Support/typora-user-images/image-20201002233038391.png)



Golang中的多协程有点类似其他语言中的多线程.

多协程和多线程:  Golang中每个goroutine(协程)默认占用内存远比Java 、C的线程少.

OS线程(操作系统线程)一般都有固定的栈内存(通常为2MB左右),一个goroutine(协程)占用内存非常小,只有2KB左右,多协程goroutine切换调度开销方面远比线程要少.

这也是为什么越来越多的大公司使用Golang的原因之一.



### 四、Goroutine的使用以及sync.WaitGroup

并行执行需求:

在主线程(可以理解成进程)中,开启一个goroutine, 该协程每隔50毫秒输出"你好golang"

在主线程中也每隔50毫秒输出"你好golang", 输出10 次后,退出程序,要求主线程和goroutine同时执行.

```go
package main
import (
"fmt"
"strconv"
"time"
)

func test() {
for i := 1; i <= 10; i++ {
fmt.Println("tesst () hello,world " + strconv.Itoa(i))
time.Sleep(time.Second)
}
}

func main() {
go test() // 开启了一个协程
for i := 1; i <= 10; i++ {
fmt.Println(" main() hello,golang" + strconv.Itoa(i))
time.Sleep(time.Second)
}
}
```

上面代码看上去没有问题,但是要注意主线程执行完毕后即使协程没有执行完毕,程序

会退出,所以需要对上面代码进行改造.



![image-20201002233335742](/Users/banna/Library/Application Support/typora-user-images/image-20201002233335742.png)



sync.WaitGroup 可以实现主线程等待协程执行完毕.

```go
package main
import (
"fmt"
"strconv"
"sync"
"time"
)

var wg sync.WaitGroup //1、定义全局的WaitGroup

func test() {
for i := 1; i <= 10; i++ {
fmt.Println("test () 你好golang " + strconv.Itoa(i))
time.Sleep(time.Millisecond * 50)
}
wg.Done() // 4、goroutine结束就登记-1
}

func main() {
wg.Add(1) //2、启动一个goroutine就登记+1
go test()
for i := 1; i <= 2; i++ {
fmt.Println(" main() 你好golang" + strconv.Itoa(i))
time.Sleep(time.Millisecond * 50)
}
wg.Wait() // 3、等待所有登记的goroutine都结束
}
```



### 五、启动多个Goroutine

在Go语言中实现并发就是这样简单,还可以启动多个goroutine.来一个例子:

(这里使用了sync.WaitGroup来实现等待goroutine执行完毕)

```go
var wg sync.WaitGroup

func hello(i int) {
defer wg.Done() // goroutine结束就登记-1
fmt.Println("Hello Goroutine!", i)
}

func main() {
for i := 0; i < 10; i++ {
wg.Add(1) // 启动一个goroutine就登记+1
go hello(i)
}
wg.Wait() // 等待所有登记的goroutine都结束
}
```



多次执行上面的代码,会发现每次打印的数字的顺序都不一致.这是因为10 个goroutine

是并发执行的,而goroutine的调度是随机的.



### 六、设置Golang并行运行的时候占用的cup数量

Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go

代码.默认值是机器上的CPU核心数.例如在一个8核心的机器上,调度器会把Go代码同

时调度到8个OS线程上.

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数.

Go1.5 版本之前,默认使用的是单核心执行. Go1.5版本之后,默认使用全部的CPU逻辑核心数.

```go
package main
import (
"fmt"
"runtime"
)

func main() {
//获取当前计算机上面的Cup个数
cpuNum := runtime.NumCPU()
fmt.Println("cpuNum=", cpuNum)
//可以自己设置使用多个cpu
runtime.GOMAXPROCS(cpuNum - 1)
fmt.Println("ok")
}
```



### 七、Goroutine统计素数

需求:  要统计1-120000 的数字中那些是素数？

##### 1、通过传统的for循环来统计

```go
func main() {
start := time.Now().Unix()
for num := 1; num <= 120000; num++ {
flag := true //假设是素数
for i := 2; i < num; i++ {
if num%i == 0 { //说明该num不是素数
flag = false
break
}
}
if flag {
// fmt.Println(num)
}
}
end := time.Now().Unix()
fmt.Println("普通的方法耗时=", end-start)
}
```



##### 2、goroutine开启多个协程统计

```go
package main

import (
"fmt"
"sync"
"time"
)

var wg sync.WaitGroup

func fn1(n int) {
for num := (n-1)*30000 + 1; num <= n*30000; num++ {
flag := true //假设是素数
for i := 2; i < num; i++ {
if num%i == 0 {
flag = false
break
}
}
if flag {
// fmt.Println(num)
}
}
wg.Done()
}

func main() {
start := time.Now().Unix()
for i := 1; i <= 4; i++ {
wg.Add(1)
go fn1(i)
}
wg.Wait()
end := time.Now().Unix()
fmt.Println("普通的方法耗时=", end-start)
}
```



问题：上面使用了goroutine已经能大大的提升新能了,但是如果想统计数据和打

印数据同时进行,这个时候如何实现呢,这个时候就可以使用管道.



### 八、Channel管道

管道是Golang在语言级别上提供的goroutine间的通讯方式,可以使用channel在

多个goroutine之间传递消息.如果说goroutine是Go程序并发的执行体,channel就是它们

之间的连接.channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制.



Golang的并发模型是CSP(Communicating Sequential Processes),提倡通过通信共享内

存而不是通过共享内存而实现通信.



Go语言中的管道(channel)是一种特殊的类型.管道像一个传送带或者队列,总是遵

循先入先出(First In First Out)的规则,保证收发数据的顺序.每一个管道都是一个具体类

型的导管,也就是声明channel的时候需要为其指定元素类型.



##### 1、channel类型

channel是一种类型,一种引用类型.声明管道类型的格式如下:

```go
var 变量chan 元素类型
举几个例子：
var ch1 chan int // 声明一个传递整型的管道
var ch2 chan bool // 声明一个传递布尔型的管道
var ch3 chan []int // 声明一个传递int 切片的管道
```



##### 2、创建channel

声明的管道需要使用make函数初始化之后才能使用.

创建channel 的格式如下:

```go
make(chan 元素类型,容量)
```



举几个例子:

```go
//创建一个能存储10个int类型数据的管道
ch1 := make(chan int, 10)
//创建一个能存储4个bool类型数据的管道
ch2 := make(chan bool, 4)
//创建一个能存储3个[]int切片类型数据的管道
ch3 := make(chan []int, 3)
```



##### 3、channel 操作

管道有发送(send)、接收(receive)和关闭(close)三种操作.

发送和接收都使用<-符号.

现在先使用以下语句定义一个管道:

```go
ch := make(chan int, 3)
```



###### 1、发送(将数据放在管道内)

将一个值发送到管道中.

```go
ch <- 10 // 把10 发送到ch 中
```



###### 2、接收(从管道内取值)

从一个管道中接收值.

```go
x := <- ch // 从ch 中接收值并赋值给变量x
<-ch // 从ch 中接收值,忽略结果
```



###### 3、关闭管道

通过调用内置的close函数来关闭管道.

```go
close(ch)
```



关于关闭管道需要注意的事情是,只有在通知接收方goroutine所有的数据都发送完毕的时

候才需要关闭管道.管道是可以被垃圾回收机制回收的,它和关闭文件是不一样的,在结束

操作之后关闭文件是必须要做的,但关闭管道不是必须的.



###### 关闭后的管道有以下特点:

​	1.对一个关闭的管道再发送值就会导致panic.

​	2.对一个关闭的管道进行接收会一直获取值直到管道为空

​	3.对一个关闭的并且没有值的管道执行接收操作会得到对应类型的零值.

​	4.关闭一个已经关闭的管道会导致panic



##### 4、管道阻塞

###### 1、无缓冲的管道:

​	如果创建管道的时候没有指定容量,那么可以叫这个管道为无缓冲的管道

​	无缓冲的管道又称为阻塞的管道.来看一下下面的代码:

```go
func main() {
ch := make(chan int)
ch <- 10
fmt.Println("发送成功")
}
```



上面这段代码能够通过编译,但是执行的时候会出现以下错误:

```go
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
D:/go_demo/demo21/07goroutine/main.go:10 +0x5b
exit status 2
```



###### 2、有缓冲的管道:

解决上面问题的方法还有一种就是使用有缓冲区的管道.可以在使用make函数初始化

管道的时候为其指定管道的容量,例如:

```go
func main() {
ch := make(chan int, 1) // 创建一个容量为1 的有缓冲区管道
ch <- 10
fmt.Println("发送成功")
}
```



只要管道的容量大于零,那么该管道就是有缓冲的管道,管道的容量表示管道中能存放元素

的数量.就像小区的快递柜只有那么多个格子,格子满了就装不下了,就阻塞了,等到别

人取走一个,快递员就能往里面放一个.



管道阻塞具体代码如下:

```go
func main() {
ch := make(chan int, 1)
ch <- 10
ch <- 12
fmt.Println("发送成功")
}
```

解决办法:

```go
func main() {
ch := make(chan int, 1)
ch <- 10 //放进去
<-ch //取走
ch <- 12 //放进去
<-ch //取走
ch <- 17 //还可以放进去
fmt.Println("发送成功")
}
```



##### 5、for range从管道循环取值

当向管道中发送完数据时,可以通过close函数来关闭管道.

当管道被关闭时,再往该管道发送值会引发panic,从该管道取值的操作会先取完管道中的

值,再然后取到的值一直都是对应类型的零值.那如何判断一个管道是否被关闭了呢？

来看下面这个例子:

```go
package main
import "fmt"

//循环遍历管道数据
func main() {
var ch1 = make(chan int, 5)
for i := 0; i < 5; i++ {
ch1 <- i + 1
}
close(ch1) //关闭管道
//使用for range遍历管道,当管道被关闭的时候就会退出for range,如果没有关闭管道
就会报个错误fatal error: all goroutines are asleep - deadlock!
//通过for range来遍历管道数据管道没有key
for val := range ch1 {
fmt.Println(val)
}
}
```



从上面的例子中看到有两种方式在接收值的时候判断该管道是否被关闭,不过通常

使用的是for range的方式.使用for range遍历管道,当管道被关闭的时候就会退出for range.



### 九、Goroutine结合Channel管道

需求1:  定义两个方法,一个方法给管道里面写数据,一个给管道里面读取数据.要求同步进行.

1、开启一个fn1的的协程给向管道inChan中写入100条数据

2、开启一个fn2的协程读取inChan中写入的数据

3、注意: fn1和fn2同时操作一个管道

4、主线程必须等待操作完成后才可以退出



```go
package main
import (
"fmt"
"sync"
"time"
)

var wg sync.WaitGroup

func fn1(intChan chan int) {
for i := 0; i < 100; i++ {
intChan <- i + 1
fmt.Println("writeData 写入数据-", i+1)
time.Sleep(time.Millisecond * 100)
}
close(intChan)
wg.Done()
}

func fn2(intChan chan int) {
for v := range intChan {
fmt.Printf("readData 读到数据=%v\n", v)
time.Sleep(time.Millisecond * 50)
}
wg.Done()
}

func main() {
allChan := make(chan int, 100)
wg.Add(1)
go fn1(allChan)
wg.Add(1)
go fn2(allChan)
wg.Wait()
fmt.Println("读取完毕...")
}
```



需求2:  goroutine结合channel实现统计1-120000的数字中那些是素数?

```go
package main
import (
"fmt"
"sync"
"time"
)

var wg sync.WaitGroup
//向intChan放入1-120000个数
func putNum(intChan chan int) {
for i := 1; i <= 1000; i++ {
intChan <- i
}
//关闭intChan
close(intChan)
wg.Done()
}

// 从intChan取出数据,并判断是否为素数,如果是,就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
for num := range intChan {
var flag bool = true
for i := 2; i < num; i++ {
if num%i == 0 { //说明该num不是素数
flag = false
break
}
}
if flag {
//将这个数就放入到primeChan
primeChan <- num
}
}
//判断关闭
exitChan <- true
wg.Done()
}

//打印素数的方法
func printPrime(primeChan chan int) {
for v := range primeChan {
fmt.Println(v)
}
wg.Done()
}

func main() {
start := time.Now().Unix()
intChan := make(chan int, 1000)
primeChan := make(chan int, 20000) //放入结果
//标识退出的管道
exitChan := make(chan bool, 8) //8个
//开启一个协程,向intChan放入1-8000个数
wg.Add(1)
go putNum(intChan)
//开启4个协程,从intChan取出数据,并判断是否为素数,如果是,就放入到primeChan
for i := 0; i < 8; i++ {
wg.Add(1)
go primeNum(intChan, primeChan, exitChan)
}
//打印素数
wg.Add(1)
go printPrime(primeChan)
//判断什么时候退出
wg.Add(1)
go func() {
for i := 0; i < 8; i++ {
<-exitChan
}
//当从exitChan取出了8个结果,就可以放心的关闭primeChan
close(primeChan)
wg.Done()
}()
wg.Wait()
end := time.Now().Unix()
fmt.Println(end - start)
fmt.Println("main 线程退出")
}
```



### 十、单向管道

有的时候会将管道作为参数在多个任务函数间传递,很多时候在不同的任务函数中

使用管道都会对其进行限制,比如限制管道在函数中只能发送或只能接收.

例如:

```go
//1. 在默认情况下下,管道是双向
//var chan1 chan int //可读可写
//2 声明为只写
var chan2 chan<- int
chan2 = make(chan int, 3)
chan2<- 20
//num := <-chan2 //error
fmt.Println("chan2=", chan2)
//3. 声明为只读
var chan3 <-chan int
num2 := <-chan3
//chan3<- 30 //err
fmt.Println("num2", num2)
```



### 十一、select多路复用

传统的方法在遍历管道时,如果不关闭会阻塞而导致deadlock,在实际开发中,可能不好确定什么关闭该管道.

也许会写出如下代码使用遍历的方式来实现:

```go
for{
// 尝试从ch1 接收值
data, ok := <-ch1
// 尝试从ch2 接收值
data, ok := <-ch2
…
}
```



这种方式虽然可以实现从多个管道接收值的需求,但是运行性能会差很多.为了应对这种场景,Go内置了select关键字,可以同时响应多个管道的操作.

select的使用类似于switch语句,它有一系列case分支和一个默认的分支.每个case会对

应一个管道的通信(接收或发送)过程.select会一直等待,直到某个case的通信操作完成

时,就会执行case分支对应的语句.具体格式如下:

```go
select{
case <-ch1:
...
case data := <-ch2:
...
case ch3<-data:
...
default:
默认操作
}
```



举个小例子来演示下select的使用:

```go
package main
import (
"fmt"
"time"
)

func main() {
//使用select可以解决从管道取数据的阻塞问题,传统的方法在遍历管道时,如果不关闭会阻塞而导致deadlock,在实际开发中,可能不好确定什么关闭该管道.
//1.定义一个管道10个数据int
intChan := make(chan int, 10)
for i := 0; i < 10; i++ {
intChan <- i
}
//2.定义一个管道5个数据string
stringChan := make(chan string, 5)
for i := 0; i < 5; i++ {
stringChan <- "hello" + fmt.Sprintf("%d", i)
}
for {
select {
//注意: 这里,如果intChan一直没有关闭,不会一直阻塞而deadlock,会自动到下一个case匹配
case v := <-intChan:
fmt.Printf("从intChan 读取的数据%d\n", v)
time.Sleep(time.Second)
case v := <-stringChan:
fmt.Printf("从stringChan 读取的数据%s\n", v)
time.Sleep(time.Second)
default:
fmt.Printf("都取不到了,不玩了, 程序员可以加入逻辑\n")
time.Sleep(time.Second)
return
}
}
}
```



使用select语句能提高代码的可读性.

​	• 可处理一个或多个channel的发送/接收操作.

​	• 如果多个case同时满足,select会随机选择一个.

​	• 对于没有case的select{}会一直等待,可用于阻塞main函数.



### 十二、Golang并发安全和锁

需求:  现在要计算1-60的各个数的阶乘,并且把各个数的阶乘放入到map中.最后显示出

来.要求使用goroutine完成.

思路

​	1.编写一个函数,来计算各个数的阶乘,并放入到map中.

​	2.启动多个协程,将统计的将结果放入到map中只使用Goroutine实现,运行的时候可能会出现资源争夺问题concurrent map writes：

```go
package main
import (
"fmt"
"sync"
_ "time"
)

var (
myMap = make(map[int]int)
wg sync.WaitGroup
)

// test函数就是计算n!,让将这个结果放入到myMap
func test(n int) {
res := 1
for i := 1; i <= n; i++ {
res *= i
}
myMap[n] = res
wg.Done()
}

func main() {
for i := 1; i <= 60; i++ {
wg.Add(1)
go test(i)
}
wg.Wait()
for i, v := range myMap {
fmt.Printf("map[%d]=%d\n", i, v)
}
}
```



##### 1、互斥锁

互斥锁是一种常用的控制共享资源访问的方法,它能够保证同时只有一个goroutine可以访

问共享资源.Go语言中使用sync包的Mutex类型来实现互斥锁.使用互斥锁来修复上面

代码的问题:

```go
package main
import (
"fmt"
"sync"
_ "time"
)

var (
myMap = make(map[int]int)
wg sync.WaitGroup
lock sync.Mutex
)

// test函数就是计算n!,让将这个结果放入到myMap
func test(n int) {
res := 1
for i := 1; i <= n; i++ {
res *= i
}
//加锁
lock.Lock()
myMap[n] = res
//解锁
lock.Unlock()
wg.Done()
}

func main() {
for i := 1; i <= 60; i++ {
wg.Add(1)
go test(i)
}
wg.Wait()
for i, v := range myMap {
fmt.Printf("map[%d]=%d\n", i, v)
}
}
```



使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区,其他的goroutine则在等

待锁;当互斥锁释放后,等待的goroutine才可以获取锁进入临界区,多个goroutine同时等

待一个锁时,唤醒的策略是随机的.

虽然使用互斥锁能解决资源争夺问题,但是并不完美,通过全局变量加锁同步来实现通讯,

并不利于多个协程对全局变量的读写操作.这个时候也可以通过另一种方式来实现上面

的功能管道(Channel).



##### 2、读写互斥锁

互斥锁是完全互斥的,但是有很多实际的场景下是读多写少的,当并发的去读取一个资

源不涉及资源修改的时候是没有必要加锁的,这种场景下使用读写锁是更好的一种选择.读

写锁在Go语言中使用sync包中的RWMutex类型.

读写锁分为两种:  读锁和写锁.当一个goroutine获取读锁之后,其他的goroutine如果是获

取读锁会继续获得锁,如果是获取写锁就会等待;当一个goroutine获取写锁之后,其他的

goroutine无论是获取读锁还是写锁都会等待.

读写锁示例:

```go
var (
x int64
wg sync.WaitGroup
lock sync.Mutex
rwlock sync.RWMutex
)

func write() {
// lock.Lock() // 加互斥锁
rwlock.Lock() // 加写锁
x = x + 1
time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
rwlock.Unlock() // 解写锁
// lock.Unlock() // 解互斥锁
wg.Done()
}

func read() {
// lock.Lock() // 加互斥锁
rwlock.RLock() // 加读锁
time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
rwlock.RUnlock() // 解读锁
// lock.Unlock() // 解互斥锁
wg.Done()
}

func main() {
start := time.Now()
for i := 0; i < 10; i++ {
wg.Add(1)
go write()
}
for i := 0; i < 1000; i++ {
wg.Add(1)
go read()
}
wg.Wait()
end := time.Now()
fmt.Println(end.Sub(start))
}
```



需要注意的是读写锁非常适合读多写少的场景,如果读和写的操作差别不大,读写锁的优势就发挥不出来.



### 十三、Goroutine Recover解决协程中出现的Panic

```go
package main
import (
"fmt"
"time"
)

//函数
func sayHello() {
for i := 0; i < 10; i++ {
time.Sleep(time.Second)
fmt.Println("hello,world")
}
}

//函数
func test() {
//这里我们可以使用defer + recover
defer func() {
//捕获test 抛出的panic
if err := recover(); err != nil {
fmt.Println("test() 发生错误", err)
}
}()
  
//定义了一个map
var myMap map[int]string
myMap[0] = "golang" //error
}

func main() {
go sayHello()
go test()
for i := 0; i < 10; i++ {
fmt.Println("main() ok=", i)
time.Sleep(time.Second)
}
}

```


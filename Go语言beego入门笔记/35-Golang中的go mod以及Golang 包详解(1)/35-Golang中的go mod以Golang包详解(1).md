一、Golang 中包的介绍和定义.......................................................................................... 1

二、Golang 包管理工具go mod........................................................................................1

三、Golang 中自定义包...................................................................................................... 3

1、定义一个包.................................................................................................................. 3

2、导入一个包.................................................................................................................. 5

四、Golang 中init()初始化函数......................................................................................... 6

五、Golang 中使用第三方包.............................................................................................. 7



### 一、Golang 中包的介绍和定义

包（package）是多个Go 源码的集合,是一种高级的代码复用方案,Go语言为提供了

很多内置包,如fmt、strconv、strings、sort、errors、time、encoding/json、os、io 等.

Golang 中的包可以分为三种:  1、系统内置包2、自定义包3、第三方包

系统内置包:  Golang语言提供的内置包,引入后可以直接使用,如fmt、strconv、strings、

sort、errors、time、encoding/json、os、io 等.

自定义包:  开发者自己写的包

第三方包:  属于自定义包的一种,需要下载安装到本地后才可以使用,如前面给大家介绍的

"github.com/shopspring/decimal"包解决float 精度丢失问题.



### 二、Golang 包管理工具go mod

在Golang1.11版本之前如果我们要自定义包的话必须把项目放在GOPATH 目录.Go1.11 版

本之后无需手动配置环境变量,使用go mod管理项目,也不需要非得把项目放到GOPATH

指定目录下,可以在你磁盘的任何位置新建一个项目, Go1.13 以后可以彻底不要GOPATH

了.



#### 1、go mod init 初始化项目

实际项目开发中首先要在项目目录中用go mod命令生成一个go.mod文件管理项目的依赖.

比如golang项目文件要放在了itying这个文件夹,这个时候需要在itying文件夹

里面使用go mod命令生成一个go.mod文件

![image-20200930211721273](/Users/banna/Library/Application Support/typora-user-images/image-20200930211721273.png)



![image-20200930211736411](/Users/banna/Library/Application Support/typora-user-images/image-20200930211736411.png)



#### 2、go mod 其他命令

![image-20200930211810842](/Users/banna/Library/Application Support/typora-user-images/image-20200930211810842.png)



download download modules to local cache (下载依赖的module 到本地cache))

edit edit go.mod from tools or scripts (编辑go.mod 文件)

graph print module requirement graph (打印模块依赖图))

init initialize new module in current directory (再当前文件夹下初始化一个新的module,创建go.mod文件))

tidy add missing and remove unused modules (增加丢失的module,去掉未用的module)

vendor make vendored copy of dependencies (将依赖复制到vendor 下)

verify verify dependencies have expected content (校验依赖检查下载的第三方库有没

有本地修改,如果有修改,则会返回非0,否则验证成功.)

why explain why packages or modules are needed (解释为什么需要依赖)



### 三、Golang 中自定义包

包(package)是多个Go源码的集合,一个包可以简单理解为一个存放多个.go 文件的文件

夹.该文件夹下面的所有go文件都要在代码的第一行添加如下代码,声明该文件归属的包.

```golang
package 包名
```



注意事项:

• 一个文件夹下面直接包含的文件只能归属一个package,同样一个package 的文件不能在多个文件夹下.

• 包名可以不和文件夹的名字一样,包名不能包含- 符号.

• 包名为main 的包为应用程序的入口包,这种包编译后会得到一个可执行文件,而编译不包含main包的源代码则不会得到可执行文件.



#### 1、定义一个包

如果想在一个包中引用另外一个包里的标识符(如变量、常量、类型、函数等)时,该标识

符必须是对外可见的(public).在Go语言中只需要将标识符的首字母大写就可以让标识

符对外可见了.

###### 1、定义一个包名为calc 的包,代码如下:

```golang
package calc
//首字母大小表示公有，首字母小写表示私有
var a = 100 //私有变量
var Age = 20 //公有变量
func Add(x, y int) int {
return x + y
}
func Sum(x, y int) int {
return x - y
}
```



###### 2、main.go 中引入这个包

访问一个包里面的公有属性方法的时候需要通过包名称.去访问

```golang
package main
import (
"fmt"
"itying/calc"
)
func main() {
c := calc.Add(10, 20)
fmt.Println(c)
}
```



#### 2、导入一个包

单行导入

单行导入的格式如下:

```golang
import "包1"
import "包2"
```



多行导入

多行导入的格式如下:

```golang
import (
"包1"
"包2"
)
```



###### 匿名导入包

如果只希望导入包,而不使用包内部的数据时,可以使用匿名导入包.具体的格式如下:

```golang
import _ "包的路径"
```

匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中.



###### 自定义包名

在导入包名的时候,还可以为导入的包设置别名.通常用于导入的包名太长或者导入的

包名冲突的情况.具体语法格式如下:

```golang
import 别名"包的路径"
```



单行引入定义别名:

```golang
import c "itying/calc"
```



多行引入定义别名:

```golang
import (
"fmt"
c "itying/calc"
)
```



![image-20200930212546087](/Users/banna/Library/Application Support/typora-user-images/image-20200930212546087.png)



### 四、Golang 中init()初始化函数

init()函数介绍

在Go 语言程序执行时导入包语句会自动触发包内部init()函数的调用.需要注意的是:  init()

函数没有参数也没有返回值. init()函数在程序运行时自动被调用执行,不能在代码中主动调用它.

包初始化执行的顺序如下图所示:

![image-20200930212632589](/Users/banna/Library/Application Support/typora-user-images/image-20200930212632589.png)



###### init()函数执行顺序

Go 语言包会从main包开始检查其导入的所有包,每个包中又可能导入了其他的包.Go编

译器由此构建出一个树状的包引用关系,再根据引用顺序决定编译顺序,依次编译这些包的代码.

在运行时,被最后导入的包会最先初始化并调用其init()函数,如下图示:

![image-20200930212730644](/Users/banna/Library/Application Support/typora-user-images/image-20200930212730644.png)



### 五、Golang 中使用第三方包

可以在https://pkg.go.dev/ 查找看常见的golang 第三方包

#### 1、找到我们需要下载安装的第三方包的地址

比如前面给大家演示的解决float 精度损失的包decimal

https://github.com/shopspring/decimal



#### 2、安装这个包

第一种方法:  go get 包名称(全局)

```golang
go get github.com/shopspring/decimal
```



第二种方法:  go mod download(全局)

```golang
go mod download
```



赖包会自动下载到$GOPATH/pkg/mod,多个项目可以共享缓存的mod,注意使用go mod

download 的时候首先需要在你的项目里面引入第三方包

第三种方法:  go mod vendor将依赖复制到当前项目的vendor下(本项目)

```golang
go mod vendor
```

将依赖复制到当前项目的vendor 下

注意:  使用go mod vendor的时候首先需要在你的项目里面引入第三方包



###### 3、看文档使用这个包

包安装完毕后就可以看文档使用这个包了...
1、Go 环境mac 环境搭建................................................................................................1

2、Go 语言开发工具Vscode 配置....................................................................................2

3、Go 语言vscode 插件安装失败解决方法......................................................................3

4、第一个Go 语言程序......................................................................................................4



### 1、Go 语言mac环境搭建

##### 1、下载安装Golang

​	Go 官网下载地址:  https://golang.org/dl/

​	Go 官方镜像站(推荐):  https://golang.google.cn/dl/



##### 2、安装软件

​	1、双击下一步下一步进行安装

![image-20201029171645175](/Users/banna/Library/Application Support/typora-user-images/image-20201029171645175.png)



​	2、验证有没有安装成功

```shell
go version
```



​	3、查看go环境

```shell
go env
```



说明:  Go1.11版本之后无需手动配置环境变量,使用go mod管理项目,也不需要非得把项

目放到GOPATH指定目录下,可以在磁盘的任何位置新建一个项目.

Go1.13以后可以彻底不要GOPATH了.



### 2、Go语言开发工具Vscode配置

##### 1、下载安装vscode

​	https://code.visualstudio.com/

##### 2、汉化vscode

![image-20201029171857054](/Users/banna/Library/Application Support/typora-user-images/image-20201029171857054.png)



##### 3、vscode中安装Go语言插件

![image-20201029171927462](/Users/banna/Library/Application Support/typora-user-images/image-20201029171927462.png)



![image-20201029171944056](/Users/banna/Library/Application Support/typora-user-images/image-20201029171944056.png)



### 3、Go 语言vscode插件安装失败解决方法

1、科学上网(搭个梯子)

2、手机开启热点,电脑连接手机热点,然后重新打开vscode下载插件

3、多试几次

![image-20201029172033137](/Users/banna/Library/Application Support/typora-user-images/image-20201029172033137.png)



### 4、第一个Go程序

```go
package main

import "fmt"

func main() {
	// fmt.Println("你好golang")
	fmt.Println("你好golang")

}
```


1、Go 语言介绍..................................................................................................................1

2、Go 语言成功的项目.......................................................................................................2

3、哪些大公司在用go 语言.................................................................................................3

4、Go 环境win 环境搭建.....................................................................................................4

5、Go 语言开发工具Vscode 配置.......................................................................................5

6、Go 语言vscode 插件安装失败解决方法..........................................................................6

7、第一个Go程序.................................................................................................................7



### 1、Go 语言介绍

Go即Golang是Google公司2009年11月正式对外公开的一门编程语言.

根据Go语言开发者自述,近10多年,从单机时代的C语言到现在互联网时代的Java,

都没有令人满意的开发语言,而C++往往给人的感觉是,花了100%的经历,却只有60%的

开发效率,产出比太低,Java和C#的哲学又来源于C++.并且,随着硬件的不断升级,这些

语言不能充分的利用硬件及CPU.因此,一门高效、简洁、开源的语言诞生了.

Go语言不仅拥有静态编译语言的安全和高性能,而且又达到了动态语言开发速度和易

维护性.有人形容Go语言：Go = C + Python,说明Go语言既有C语言程序的运行速度,又

能达到Python语言的快速开发.

Go语言是非常有潜力的语言,是因为它的应用场景是目前互联网非常热门的几个领域,

比如WEB 开发、区块链开发、大型游戏服务端开发、分布式/云计算开发.国内比较知名的

B站就是用Go语言开发的,像Goggle、阿里、京东、百度、腾讯、小米、360的很多应用

也是使用Go语言开发的.



### 2、Go语言成功的项目

nsq：bitly开源的消息队列系统,性能非常高,目前他们每天处理数十亿条的消息

docker：基于lxc的一个虚拟打包工具,能够实现PAAS平台的组建

packer：用来生成不同平台的镜像文件,例如VM、vbox、AWS 等,作者是vagrant的作者

skynet：分布式调度框架

Doozer：分布式同步工具,类似ZooKeeper

Heka：mazila开源的日志处理系统

cbfs：couchbase开源的分布式文件系统

tsuru：开源的PAAS平台,和SAE实现的功能一模一样

groupcache：memcahe作者写的用于Google下载系统的缓存系统

god：类似redis的缓存系统,但是支持分布式和扩展性

gor：网络流量抓包和重放工具



### 3、各大公司Go开源项目地址

##### Google

https://github.com/google/



##### Facebook

https://github.com/facebookgo



##### 腾讯

http://www.infoq.com/cn/articles/tencent-millions-scale-docker-applicationpractice



##### 百度

http://www.infoq.com/cn/presentations/application-of-golang-in-baidu-frontend



##### 阿里

暂无



##### 京东

暂无



##### 小米

http://open-falcon.com/



##### 360

https://github.com/Qihoo360/poseidon

https://github.com/Qihoo360/wayne



### 4、Go 环境win环境搭建

##### 1、下载安装Golang

​	Go 官网下载地址:  https://golang.org/dl/

​	Go 官方镜像站(推荐):  https://golang.google.cn/dl/

##### 2、安装软件

​	1、双击下一步下一步进行安装

![image-20201029165737891](/Users/banna/Library/Application Support/typora-user-images/image-20201029165737891.png)

​	2、验证有没有安装成功

```shell
go version
```

![image-20201029165822274](/Users/banna/Library/Application Support/typora-user-images/image-20201029165822274.png)

​	3、查看go环境

```shell
go env
```

![image-20201029170013808](/Users/banna/Library/Application Support/typora-user-images/image-20201029170013808.png)



说明:  Go1.11版本之后无需手动配置环境变量,使用go mod管理项目,也不需要非得把项

目放到GOPATH指定目录下,可以在磁盘的任何位置新建一个项目.

Go1.13以后可以彻底不要GOPATH了.



### 5、Go语言开发工具Vscode配置

##### 1、下载安装vscode

​	https://code.visualstudio.com/

##### 2、汉化vscode

![image-20201029170202595](/Users/banna/Library/Application Support/typora-user-images/image-20201029170202595.png)



##### 3、vscode中安装Go语言插件

![image-20201029170239454](/Users/banna/Library/Application Support/typora-user-images/image-20201029170239454.png)



![image-20201029170255812](/Users/banna/Library/Application Support/typora-user-images/image-20201029170255812.png)



### 6、Go语言vscode插件安装失败解决方法

1、科学上网(搭个梯子)

2、手机开启热点电,脑连接手机热点,然后重新打开vscode下载插件

3、多试几次

![image-20201029170416785](/Users/banna/Library/Application Support/typora-user-images/image-20201029170416785.png)



### 7、第一个Go程序
```go
package main

import "fmt"

func main() {
	// fmt.Println("你好golang")
	fmt.Println("你好golang")

}
```
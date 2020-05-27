package add

/*这里只是引用这个包做初始化,不使用这个函数的任何变量等*/
import (
	_ "day2/example3/test"
)

/*全局变量*/
var Name string = "xxxxxx"
var Age int = 100

/*包的init函数是在main函数之前执行的*/
func init() {
	Name = "hello world"
	Age = 10
}

/*
第一步是初始化全局变量
第二部是初始化init函数,不需要自己调用的;只用写实现init函数,然后go的运行时库它会自动调用init函数来初始化
第三部是执行main函数
备注:
	所以在main函数中调用,add自定义函数输出的值是init函数中的值
*/

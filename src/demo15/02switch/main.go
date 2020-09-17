package main

import "fmt"

func main() {
	/*
		switch exprression {
		case condition:
		}
	*/

	//练习: 判断文件类型,如果后缀名是.html 输入text/html,如果后缀名.css 输出text/css ,如果后缀名是.js 输出

	// var ext = ".html"
	// if ext == ".html" {
	// 	fmt.Println("tesxt/html")
	// } else if ext == ".css" {
	// 	fmt.Println("text/css")
	// } else if ext == ".js" {
	// 	fmt.Println("text/javascript")
	// } else {
	// 	fmt.Println("找不到此后缀")
	// }

	//1.switch case的基本使用
	// var extname = ".html"
	// switch extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// case ".js":
	// 	fmt.Println("text/javascript")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }

	//2.switch case的另一种写法
	// switch extname := ".html"; extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// case ".js":
	// 	fmt.Println("text/javascript")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }

	//3.一个分支可以有多个值,多个case值中间使用英文逗号分隔
	//判断一个数是不是偶数
	// var n = 5
	// switch n {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("奇数")
	// 	break //golang中break可写也可以不写
	// case 2, 4, 6, 8, 10:
	// 	fmt.Println("偶数")
	// 	break

	// }

	// var score = "A"
	// switch score {
	// case "A", "B", "C":
	// 	fmt.Println("及格")
	// case "D":
	// 	fmt.Println("不及格")
	// }

	// switch score := "A"; score {
	// case "A", "B", "C":
	// 	fmt.Println("及格")
	// case "D":
	// 	fmt.Println("不及格")
	// }

	//4.分支还可以使用表达式,这时候switch语句后面不需要跟判断变量,例如:
	// var age = 30
	// switch { //表达式(条件判断),switch后面不需要跟变量
	// case age < 24:
	// 	fmt.Println("好好学习")
	// case age >= 24 && age <= 60:
	// 	fmt.Println("好好赚钱")
	// case age > 60:
	// 	fmt.Println("注意身体")
	// default:
	// 	fmt.Println("输入错误")
	// }

	//5.switch 的穿透fallthrought
	//fallthrought语法可以执行满足条件的case下一个case,是为了兼容C语言中的case设计的
	var age = 30
	switch { //表达式,switch后面不需要跟变量
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age <= 60:
		fmt.Println("好好赚钱")
		fallthrough //一次只能穿透一层,只会穿透和fallthrouht紧挨着的case
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}
}

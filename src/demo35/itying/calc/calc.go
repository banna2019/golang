package calc

var aaa = "私有变量" //首字母小写表示私有
var Age = 20

func Add(x, y int) int { //公有方法
	return x + y

}

func Sub(x, y int) int { //公有方法
	return x - y
}

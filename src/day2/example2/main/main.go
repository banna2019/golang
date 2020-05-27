package main

/*原始应用包*/
// import (
// 	"day2/example2/add"
// 	"fmt"
// )

/*包的别名引用*/
import (
	a "day2/example2/add"
	"fmt"
)

/*第一种写法*/
// func main() {
// 	fmt.Println("Name=", add.Name)
// 	fmt.Println("Age=", add.Age)
// }

/*第二种写法*/
// func main() {
// 	add.Test()
// 	fmt.Println("Name=", add.Name)
// 	fmt.Println("Age=", add.Age)
// }

/*包的别名引用*/
func main() {
	a.Test()
	fmt.Println("Name=", a.Name)
	fmt.Println("Age=", a.Age)
}

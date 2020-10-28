package main

func main() {

	// /1、删除一个文件(当删除的目标不存在是会报: "remove aaa.txt: no such file or directory")
	// err := os.Remove("aaa.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//2、删除一个目录(当删除目标不在时候,汇报: "remove ./aaa: no such file or directory")
	// err := os.Remove("./aaa")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 3、一次删除多个文件(当删除目标不存在的时候,不会做任何操作)
	// err := os.RemoveAll("dir1")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

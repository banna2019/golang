package test

import (
	"fmt"
)

var Name string = "This is in test package"
var Age int = 1000

func init() {
	fmt.Println("This is a test")
	fmt.Println("test.package.Name=", Name)
	fmt.Println("test.package.Age=", Age)

	Age = 10
	fmt.Println("test.package.Age=", Age)
}

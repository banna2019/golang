package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is %s\n", goos)
	var path string = os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "    hello world abc    \n"
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println("Replace:", result)

	count := strings.Count(str, "1")
	fmt.Println("Count:", count)

	result = strings.Repeat(str, 3)
	fmt.Println("Repeat:", result)

	result = strings.ToLower(str)
	fmt.Println("Lower:", result)

	result = strings.ToUpper(str)
	fmt.Println("Upper:", result)

	result = strings.TrimSpace(str)
	fmt.Println("TrimSpace:", result)

	result = strings.Trim(str, "\n\r")
	fmt.Println("trim:", result)

	result = strings.TrimLeft(str, "\n\r")
	fmt.Println("TrimLeft:", "\n\r")

	result = strings.TrimRight(str, "\n\r")
	fmt.Println("TrimRight:", result)

	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(str, "l")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "l")
	fmt.Println("Join:", str2)

	str2 = strconv.Itoa(1000)
	fmt.Println("Itoa:", str2)

	number, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("can not convert ot int,", err)
		return
	}
	fmt.Println("number:", number)
}

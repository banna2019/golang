package main

import (
	"fmt"
	"time"
)

//1.打印当前时间

func main() {

	timeObj := time.Now()
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	// fmt.Printf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

	//注意: %02d 中的2 表示宽度,如果整数不够2 列就补上0
}

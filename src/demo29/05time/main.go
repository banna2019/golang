package main

import (
	"fmt"
	"time"
)

//日期字符串转换成时间戳
func main() {
	var str = "2020-09-18 05:33:45"
	var tmp = "2006-01-02 15:04:05"

	timeObj, _ := time.ParseInLocation(tmp, str, time.Local)
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

}

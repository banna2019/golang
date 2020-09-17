package main

import (
	"fmt"
	"time"
)

//时间戳转换成日期字符串
func main() {
	//unixtime: 1600377919

	// var unixTime int64 = 1600377919
	// timeObj := time.Unix(unixTime, 0)
	// fmt.Println(timeObj.Format("2006-01-02 15:04:05"))

	var unixTime = 1600377919
	timeObj := time.Unix(int64(unixTime), 0)
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))

	//时间戳格式化成时间时间字符串
	timeObj1 := time.Now()
	timeNow := timeObj1.Unix()
	fmt.Println(timeNow)

	timeTime := time.Unix(int64(timeNow), 0)
	var unixNowTime = timeTime.Format("2006-01-02 15:04:05")
	fmt.Println(unixNowTime)

}

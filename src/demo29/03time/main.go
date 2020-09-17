package main

import (
	"fmt"
	"time"
)

//获取当前时间
/*
	时间戳是自1970年1月1日(08:00:00GMT)至当前时间的总毫秒数.它也被称为Unix时间戳(UnixTimestamp).
*/
func main() {
	timeObj := time.Now()

	unixtime := timeObj.Unix() //获取当前时间戳(自1970年1月1日(08:00:00GMT)至当前时间的总毫秒数)

	fmt.Println("当前时间戳: ", unixtime)

	unixNatime := timeObj.UnixNano() //纳秒时间戳
	fmt.Println("当前时间戳: ", unixNatime)
}

package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	// var num1 float64 = 3.1
	// var num2 float64 = 4.2
	// fmt.Println(num1 + num2) //7.300000000000001

	/*
		//第三方decimal浮点数精度处理
		//加
		var num1 float64 = 3.1
		var num2 float64 = 4.2
		d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
		fmt.Println(d1)

		//相减
		m1 := 8.2
		m2 := 3.8
		m3 := decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))
		fmt.Println(m3)

		//减法 Sub,乘法 Mul,除法 Div;用法均与上述类似

		c1 := 1.2
		c2 := 3.8
		c3 := decimal.NewFromFloat(c1).Mul(decimal.NewFromFloat(c2))
		fmt.Println(c3)

		s1 := 19.9
		s2 := 0.9
		s3 := decimal.NewFromFloat(s1).Div(decimal.NewFromFloat(s2))
		fmt.Println(s3)
	*/

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

	value := gjson.Get(json, "name.last")
	fmt.Println(value.String())

}

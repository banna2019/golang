package main

func main() {
	/*
		//写一个程序,统计一个字符串中每个单词出现的次数.比如: "how do you do"中how=1 do=2 you=1

		//this is golang

		var str = "how do you do"
		var strSlice = strings.Split(str, " ")
		// fmt.Println(strSlice)

		var strMap = make(map[string]int)
		for _, v := range strSlice {
			strMap[v]++
		}
		// fmt.Println(strMap)
		for k, v := range strMap {
			fmt.Printf("key:%v value:%v\n", k, v)
		}
	*/

	/*
		s1 := "how do you do"
		s3 := strings.Split(s1, " ")
		m1 := make(map[string]int)
		for _, v := range s3 {
			_, ok := m1[v]
			if !ok {
				m1[v] = 1
			} else {
				m1[v]++
			}
		}
		fmt.Println(m1)
	*/

	/*
		s1 := "how do you do"
		s3 := strings.Split(s1, " ")
		how := 0
		do := 0
		you := 0
		for _, v := range s3 {
			switch v {
			case "how":
				how++
			case "do":
				do++
			case "you":
				you++

			}
		}
		fmt.Printf("how=%d\n", how)
		fmt.Printf("do=%d\n", do)
		fmt.Printf("you=%d\n", you)
	*/

	/*
		a := "how do you do"
		word := make([]string, 4)
		fmt.Println(word)
		couner := 0
		how := 0
		do := 0
		you := 0
		for i := range a {
			value := fmt.Sprintf("%c", a[i])
			if value != " " {
				word[couner] += value
			} else {
				couner++
			}
		}
		for _, v := range word {
			switch v {
			case "how":
				how++
			case "do":
				do++
			case "you":
				you++

			}
		}
		fmt.Println(word)
		fmt.Printf("how=%d\t", how)
		fmt.Printf("do=%d\t", do)
		fmt.Printf("you=%d\t", you)
	*/
}

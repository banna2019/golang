package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 43, 3, 78, 234, 975}
	sort.Ints(a[:])

	fmt.Println(a)
}

func testString() {
	var a = [...]string{"abdb", "efg", "A", "eee"}
	sort.Strings(a[:])

	fmt.Println(a)
}

func testfloat() {
	var a = [...]float64{2.3, 0.8, 28.2, 3258923.345, 0.6}
	sort.Float64s(a[:])

	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1, 43, 3, 78, 234, 975}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 3) //这里的查找默认是排序后的位置

	fmt.Println(a)
	fmt.Println(index)
}

func testStringSearch() {
	var a = [...]string{"abdb", "efg", "A", "eee"}
	sort.Strings(a[:])
	index := sort.SearchStrings(a[:], "A")
	fmt.Println(a)
	fmt.Println(index)
}

func testfloatSearch() {
	var a = [...]float64{2.3, 0.8, 28.2, 3258923.345, 0.6}
	sort.Float64s(a[:])
	index := sort.SearchFloat64s(a[:], 0.6)
	fmt.Println(a)
	fmt.Println(index)
}

func main() {
	// testIntSort()
	// testString()
	// testfloat()
	testIntSearch()
	testStringSearch()
	testfloatSearch()
}

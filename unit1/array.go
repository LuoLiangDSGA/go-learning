package main

import "fmt"

func array() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The First element is %d\n", arr[0])
}

func array1() {
	a := [3]int{1, 2, 3}
	// go会自动统计元素的个数
	b := [...]int{1, 2, 3}
	c := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	// create slice 可增长的数组
	s1 := a[2:4]
}

func main() {
	array()
}

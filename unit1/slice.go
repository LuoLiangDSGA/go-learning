package main

import "fmt"

type Animal struct {
	name string
	age  int8
}

func slice1() {
	arr := [3]Animal{{"a", 1}, {"b", 2}}
	fmt.Println(arr)
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	slice := ints[:5]
	fmt.Println(slice)
	a := ints[0:3]
	b := ints[3:]
	fmt.Println(a, b)
	b[0] = 1
	fmt.Println(a, b)
	fmt.Printf("len=%d, cap=%d, %v\n", len(b), cap(b), b)
	fmt.Println(ints) // 改变切片的值会对原数组有影响
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

func slice2() {
	// make创建切片
	a := make([]int, 5) // len(a) = 5
	fmt.Println(a, len(a), cap(a))
	b := make([]int, 5, 10)
	fmt.Println(b, len(b), cap(b))
	c := a[:2]
	fmt.Println(c, len(c), cap(c))
	matrix := [][]string{
		{"x", "y"},
		{"x", "y"},
		{"x", "y"},
	}
	fmt.Println(matrix)
}

func slice3() {
	// append函数应用
	var arr []int
	printSlice(arr)
	arr = append(arr, 0)
	printSlice(arr)
	arr = append(arr, 1)
	printSlice(arr)
	arr = append(arr, 2, 3, 4)
	printSlice(arr)
	pow := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for v := range pow {
		fmt.Printf("%d\n", v)
	}
}

//实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。
//当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。
//
//图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。
//
//（提示：需要使用循环来分配 [][]uint8 中的每个 []uint8；请使用 uint8(intValue) 在类型之间转换；你可能会用到 math 包中的函数。）

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//切片
func main() {
	slice1()
	slice2()
	slice3()
}

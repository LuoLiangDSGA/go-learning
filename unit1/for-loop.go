/**
*For-loop

*1. 创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。

*2. 用 goto 改写 1 的循环。关键字 for 不可使用。

*3. 再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
 */
package main

import "fmt"

func forloop1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("for result:%d\n", i)
	}
}

func forloop2() {
	i := 0
Here:
	fmt.Printf("goto result:%d\n", i)
	i++
	if i < 10 {
		goto Here
	}
}

func forloop3() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(array); i++ {
		fmt.Printf("array result:%d\n", array[i])
	}
}

func main() {
	forloop1()
	forloop2()
	forloop3()
}
package main

import ("fmt")

func f1() {
	defer fmt.Println("World")

	fmt.Printf("Hello ")
}

func f2() {
	fmt.Println("start.")

	for i := 1; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("done.")
}
//defer 语句会将函数推迟到外层函数返回之后执行。
//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
func main() {
	f1()
	f2()
}
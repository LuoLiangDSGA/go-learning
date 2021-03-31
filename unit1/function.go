package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	// 闭包 引用了函数体之外的的变量sum
	return func(x int) int {
		sum += x
		return sum
	}
}

// 返回一个“返回int的函数”
//`(0, 1, 1, 2, 3, 5, ...)`。
func fibonacci() func() int {
	a := 0
	b := 1
	i := 0
	return func() int {
		tmp := a + b
		if i == 0 {
			i++
			return 0
		}
		if i > 1 {
			a = b
			b = tmp
		}
		i++
		return tmp
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	a, b := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i), b(-2*i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

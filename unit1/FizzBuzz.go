package main

import "fmt"

func main() {
	fmt.Println("====solution1====")
	solution1()
	fmt.Println("====solution2====")
	solution2()
}

func solution1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i%5 == 0 {
			fmt.Printf("Buzz\n")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Printf("Fizz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}

func solution2() {
	//定义常量
	const (
		Fizz = 3
		Buzz = 5
	)

	var f bool
	for i := 1; i <= 100; i++ {
		f = false
		if i%Fizz == 0 {
			fmt.Printf("Fizz")
			f = true
		}
		if i%Buzz == 0 {
			fmt.Printf("Buzz")
			f = true
		}
		if !f {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}

}

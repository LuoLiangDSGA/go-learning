package main

import "fmt"

func main() {
	// solution1()
	solution2()
}

func solution1() {
	str := "A"
	for i := 1; i <= 100; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(str)
		}
		fmt.Println()
	}
}

func solution2() {
	str := "A"
	for i := 1; i <= 100; i++ {
		fmt.Printf("%s\n", str)
		str += "A"
	}
}
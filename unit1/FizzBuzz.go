package main

import "fmt"

func main() {
	solution1();
}

func solution1() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i % 5 ==0 {
			fmt.Printf("Buzz\n")
		} else if i % 3 == 0 && i % 5 == 0{
			fmt.Printf("Fizz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}


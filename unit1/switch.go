package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch1() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
		fallthrough // 会自动break，除非加上这个关键字
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch2() {
	today := time.Now().Weekday()
	switch day := time.Saturday; day {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func switch3() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning")
	case time.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}

func main() {
	switch1()
	switch2()
	switch3()
}

package main

import (
	"fmt"
)

func pointer1() {
	i, j := 10, 20
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 10
	fmt.Println(j)
}

func main() {
	pointer1()
}

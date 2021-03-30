package main

import (
	"fmt"
)

// 结构体
type Person struct {
	name string
	age  int8
}

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

//结构体练习
func struct1() {
	fmt.Println(Person{"fantasy", 1})
	v := Person{"fantasy", 1}
	v.name = "yehuimei"
	fmt.Println(v)
	fmt.Println(Person{})
	p := &v
	fmt.Println(*p)
	(*p).name = "jay chou" // 也可以通过p.name = "" 来设置值
	fmt.Println(v)
}

func main() {
	pointer1()
	struct1()
}

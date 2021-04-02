package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// golang中的方法练习，方法和函数最大的区别就是方法需要指定接受者类型
func main() {
	v := Vertex{3, 4}
	v.scale(10) //如果这里去掉*那么将不会修改原值，方法内是副本的操作
	fmt.Println(v.Abs())
	//scale(v, 10) 函数中的参数如果是指针类型，这里传普通类型，会编译报错
	scale(&v, 10)
	fmt.Println(v)
}

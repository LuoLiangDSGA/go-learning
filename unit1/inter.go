package main

import "fmt"

type Person struct {
	name string
	age  int8
}

type Stringer interface {
	String() string
}

//方法名和返回类型和接口中的方法一样，就实现了这个接口，但是如果有多个接口就需要把方法都实现之后才算实现了这个接口
func (p Person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

type MyFloat float64

func printString(s Stringer) {
	fmt.Println(s.String())
}

//这里只实现了带*号的指针参数，传入普通参数的话编译会报错
func (myFloat *MyFloat) String() string {
	return fmt.Sprintf("the value is %d", myFloat)
}

//接口练习
func main() {
	var stringer Stringer = Person{"a", 10}
	printString(stringer)
	var float = MyFloat(0.2)
	printString(&float)
}

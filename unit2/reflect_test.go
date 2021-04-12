package unit2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_jsonToStruct(t *testing.T) {

}

func Test_reflect(t *testing.T) {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	i1 := iv.Interface().(int)
	fmt.Println(iv, it, i1)
}

func Test_updateField(t *testing.T) {
	p := Person{Name: "s", Age: 1}
	v := reflect.ValueOf(&p) // ValueOf的参数这里需要指定为指针
	v.Elem().Field(0).SetString("sss")
	fmt.Println(p)
}

type Person struct {
	Name string
	Age  int
}

package unit2

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_jsonToStruct(t *testing.T) {
	p := Person{Name: "s", Age: 1}
	jsonStr, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonStr))
	}
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)
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
	fmt.Println(v.Kind())
	pt := reflect.TypeOf(p) //反射获取类型
	// 遍历字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}
	// 遍历方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}
	s := "范特西"
	//b := []byte(s)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))
	fmt.Println(b1)
}

// ``是struct tag的用法，用于把结构体中的字段名和json中的字段名做转换
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

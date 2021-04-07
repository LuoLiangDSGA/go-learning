package unit2

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go fmt.Println("hello world")
	fmt.Println("我是好人")
	time.Sleep(time.Second * 1)
}

func TestChannel(t *testing.T) {
	ch := make(chan string)
	go func() {
		fmt.Println("fantasy")
		ch <- "goroutine完成" //在右边，代表发送
	}()
	fmt.Println("i am main goroutine")
	v := <-ch
	fmt.Println("接受chan中的值为：", v)
}


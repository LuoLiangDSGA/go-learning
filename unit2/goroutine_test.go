package unit2

import (
	"fmt"
	"math/rand"
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

func TestSingleTest(t *testing.T) {
	SingleTest(make(chan int))
}

func SingleTest(out chan<- int) {
	out <- 1
	//i := <- out 编译失败，函数声明只能发送
	//fmt.Println(i)
}

func TestConcurrentDownload(t *testing.T) {
	first := make(chan string)
	second := make(chan string)
	third := make(chan string)
	go func() {
		first <- downloadFile("first file")
	}()
	go func() {
		second <- downloadFile("second file")
	}()
	go func() {
		third <- downloadFile("third file")
	}()
	select {
	case p := <-first:
		fmt.Println(p)
	case p := <-second:
		fmt.Println(p)
	case p := <-third:
		fmt.Println(p)
	}
}

func downloadFile(s string) string {
	x := rand.Intn(3)
	time.Sleep(time.Duration(x) * time.Second)
	return s + ":filePath"
}

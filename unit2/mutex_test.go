package unit2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	sum   = 0
	mutex sync.RWMutex
)

func Test_concurrentAdd(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go add(10)
	}
	for i := 0; i < 10; i++ {
		go fmt.Println("和为：", readSum())
	}
	// 防止提前退出
	time.Sleep(2 * time.Second)
}

func Test_waitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(110)
	for i := 0; i < 100; i++ { //100协程
		go func() {
			defer wg.Done()
			add1(10)
		}()
	}
	for i := 0; i < 10; i++ { //10协程
		go func() {
			defer wg.Done()
			fmt.Println("和为：", readSum1())
		}()
	}
	// 防止提前退出
	wg.Wait()
}
func Test_doOnce(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	// 创建大小为10的channel
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<- done
	}
}
func add1(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
func readSum1() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}
func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

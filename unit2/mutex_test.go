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
		<-done
	}
}

// cond的使用
func Test_cond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{}) //初始化一个条件
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位...")
			cond.L.Lock()
			cond.Wait() //等待枪响   这里wait方法注意需要加锁
			fmt.Println(num, "号开始跑！！时间：", time.Now().String())
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast()
	}()
	wg.Wait()
	// golang中的cond，wait, signal, broadcast就像Java中的wait，notify，notifyAll
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

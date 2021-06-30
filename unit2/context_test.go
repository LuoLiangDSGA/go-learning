package unit2

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_context(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog(ctx, "监控狗1")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "监控狗2")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "监控狗3")
	}()
	// 监控5s
	time.Sleep(5 * time.Second)
	stop()
	// 防止主协程退出
	wg.Wait()
}

func watchDog(ctx context.Context, s string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(s, " 的监控停止指令已收到")
			return
		default:
			fmt.Println(s, "监控中")
		}
		time.Sleep(time.Second)
	}
}

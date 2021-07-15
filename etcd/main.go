package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	rootBackground := context.Background()
	// 客户端初始化
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
	if cli == nil || err == context.DeadlineExceeded {
		// 处理错误
		fmt.Println(err)
		panic("invalid connection!")
	}
	defer cli.Close()
	//初始化kv
	kvc := clientv3.NewKV(cli)
	//获取值
	//getVal(rootBackground, err, kvc)
	//设置值
	//putVal(rootBackground, err, kvc)
	//续租测试
	leaseKV(cli, rootBackground, kvc)
	// 循环模拟kv变化
	//watchKV(kvc, err, cli)

}

func watchKV(kvc clientv3.KV, err error, cli *clientv3.Client) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("key: /aa/bb, put val cc")
			kvc.Put(context.TODO(), "/aa/bb", "cc")
			fmt.Println("delete key /aa/bb")
			kvc.Delete(context.TODO(), "/aa/bb")
			time.Sleep(1 * time.Second)
		}
	}()
	getResp, err := kvc.Get(context.TODO(), "/aa/bb")
	if err != nil {
		fmt.Println(err)
		return
	}
	// key存在
	if len(getResp.Kvs) > 0 {
		fmt.Printf("current value is %s", string(getResp.Kvs[0].Value))
	}
	watchRevision := getResp.Header.Revision + 1
	watcher := clientv3.NewWatcher(cli)
	fmt.Println("listen backwards from the current version：", watchRevision)
	watchChannel := watcher.Watch(context.TODO(), "/aa/bb", clientv3.WithRev(watchRevision))
	// 处理kv变化
	for watchResp := range watchChannel {
		for _, event := range watchResp.Events {
			switch event.Type {
			case 0:
				fmt.Println("update->", string(event.Kv.Value), "revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case 1:
				fmt.Println("delete->", "revision：", event.Kv.ModRevision)
			}
		}
	}
}

func leaseKV(cli *clientv3.Client, rootBackground context.Context, kvc clientv3.KV) {
	lease := clientv3.NewLease(cli)
	ctx, _ := context.WithTimeout(rootBackground, time.Duration(100)*time.Second)
	// 创建一个租约
	grant, _ := lease.Grant(ctx, 2)
	leaseId := grant.ID
	if _, err := kvc.Put(ctx, "/aa/bb", "cc", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	// 自动续约
	aliveCh, err := lease.KeepAlive(context.TODO(), leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			select {
			case keepResp := <-aliveCh:
				if keepResp == nil {
					fmt.Println("lease is expired")
					goto End
				} else {
					fmt.Println("receive auto lease ack：", keepResp.ID)
				}
			}
		}
	End:
	}()
	for {
		getResp, err := kvc.Get(ctx, "/aa/bb")
		if err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv expired")
			break
		}
		fmt.Println("haven't expired")
		// 延迟一秒执行
		time.Sleep(1 * time.Second)
	}
}

func putVal(rootBackground context.Context, err error, kvc clientv3.KV) {
	uuid := uuid.New().String()
	fmt.Printf("new value is :%s\r\n", uuid)
	ctx, cancelFunc2 := context.WithTimeout(rootBackground, time.Duration(2)*time.Second)
	_, err = kvc.Put(ctx, "cc", uuid)
	if delRes, err := kvc.Delete(ctx, "cc"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("delete %s for %t\n", "cc", delRes.Deleted > 0)
	}
	cancelFunc2()
	if err != nil {
		fmt.Println(err)
	}
}

func getVal(rootBackground context.Context, err error, kvc clientv3.KV) {
	ctx, cancelFunc := context.WithTimeout(rootBackground, time.Duration(2)*time.Second)
	response, err := kvc.Get(ctx, "cc")
	cancelFunc()
	if err != nil {
		fmt.Println(err)
	}
	kvs := response.Kvs
	if len(kvs) > 0 {
		fmt.Printf("last value is :%s\r\n", string(kvs[0].Value))
	} else {
		fmt.Printf("empty key for %s\n", "cc")
	}
}

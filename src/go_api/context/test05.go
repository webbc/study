package main

import (
	"time"
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")

	ctx2, cancel2 := context.WithCancel(ctx)
	go watch(ctx2, "【监控2】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控2停止")
	cancel2()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控1停止")
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

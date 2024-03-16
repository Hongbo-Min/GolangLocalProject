package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	// time.Sleep(5 * time.Second) // 运行一段时间后退出
	sig := make(chan os.Signal, 1) // 优雅退出，坚挺ctrl+c信号
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

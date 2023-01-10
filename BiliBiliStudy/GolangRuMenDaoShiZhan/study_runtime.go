package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("%d s: %v\n", i, s)
	}
}

func show2(s string) {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			fmt.Println("bye bye...")
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", i)
	}
}

func PrintA() {
	for i := 0; i < 10; i++ {
		fmt.Println("A: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func PrintB() {
	for i := 0; i < 10; i++ {
		fmt.Println("B: ", i)
	}
}

func main() {
	/* go show("java")
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让给其他子协程来执行任务
		fmt.Println("golang")
	} */
	/* go show2("java")
	time.Sleep(time.Second) */
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go PrintA()
	go PrintB()

	time.Sleep(time.Second * 1)
}

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/* var i = 100
var lock sync.Mutex

func add() {
	lock.Lock()
	i++
	lock.Unlock()
}

func sub() {
	lock.Lock()
	i--
	lock.Unlock()
} */

var i int32 = 100

// 原则操作 cas compare and swap 比较之后交换
func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("i: %v\n", i)
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Millisecond * 100)
	values <- value
}

// 通道中无数据，但执行读操作 会引发死锁
func ReadNoDataFromNoBufch() {
	noBufch := make(chan int)
	<-noBufch
	fmt.Println("read from no buffer channel success")
	// output: fatal error: all goroutines are asleep - deadlock!

}

// 通道中无数据，向通道中写入，但是无协程读取
func WriteNoBufch() {
	noBufch := make(chan int)
	noBufch <- 1
	fmt.Println("write success no block")
	// output: fatal error: all goroutines are asleep - deadlock!

}

// 通道的缓存无数据，但执行读通道
func ReadNoDataFromBufch() {
	bufch := make(chan int, 1)
	<-bufch
	fmt.Println("read from no buffer channel success")
	// output: fatal error: all goroutines are asleep - deadlock!

}

// 通道的缓存已经满了，向通道写数据，但无协程读
func WriteBufChButFull() {
	bufch := make(chan int, 1)
	bufch <- 100
	bufch <- 1
	fmt.Println("write success no block")
	// output: fatal error: all goroutines are asleep - deadlock!

}

// 从无缓冲通道中读取数据
func ReadNoDataFromNoBufChWithSelect() {
	bufch := make(chan int)
	if v, err := ReadWithSelect(bufch); err != nil {
		fmt.Println("read no data from no bufch with select err: ", err)
	} else {
		fmt.Printf("read: %d\n", v)
	}
	// output: err:  channel has no data

}

// 从有缓冲通道中读取数据
func ReadNoDataFromBufChWithSelect() {
	bufch := make(chan int, 1)
	if v, err := ReadWithSelect(bufch); err != nil {
		fmt.Println("read no data from bufch with select err: ", err)
	} else {
		fmt.Printf("read: %d\n", v)
	}
}

// 通过select读取通道
func ReadWithSelect(ch chan int) (res int, err error) {
	// 使用定时器替代default，可以将通道读取数据的容忍时间变为500ms，如果依然无法读写，就即刻返回
	timeout := time.NewTimer(time.Microsecond * 500)
	select {
	case res = <-ch:
		return
	case <-timeout.C:
		return 0, errors.New("read time out")
	}
}

// 无缓冲通道写
func WriteNoBufchWithSelect() {
	ch := make(chan int)
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println("write no bufch with select err: ", err)
	} else {
		fmt.Println("write success")
	}
}

func WriteBufChButFullWithSelect() {
	ch := make(chan int, 1)
	ch <- 100
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println("write bufch with select err: ", err)
	} else {
		fmt.Println("write success")
	}
}

// 通过select写通道
func WriteChWithSelect(ch chan int) error {
	timeout := time.NewTimer(time.Microsecond * 500)
	select {
	case ch <- 1:
		return nil
	case <-timeout.C:
		return errors.New("write time out")
	}
}

func main() {
	/* go send()
	value := <-values
	fmt.Printf("value: %v\n", value) */

	/* ReadNoDataFromNoBufch()
	WriteNoBufch()
	ReadNoDataFromBufch()
	WriteBufChButFull() */
	ReadNoDataFromNoBufChWithSelect()
	ReadNoDataFromBufChWithSelect()
	WriteNoBufchWithSelect()
	WriteBufChButFullWithSelect()
}

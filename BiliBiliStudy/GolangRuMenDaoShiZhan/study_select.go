package main

import (
	"fmt"
	"time"
)

var chInt = make(chan int, 0)
var chStr = make(chan string)

func main() {
	go func() {
		chInt <- 100
		chStr <- "hi"
		close(chInt)
		close(chStr)
	}()

	for {
		select {
		case r := <-chInt:
			fmt.Printf("chInt r: %v\n", r)
		case r := <-chStr:
			fmt.Printf("chStr r: %v\n", r)
		default:
			fmt.Println("default...")
		}

		time.Sleep(time.Second * 1)
	}
}

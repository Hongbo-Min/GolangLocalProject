package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Add(-1) // wg.Done() 功能相同
	fmt.Printf("hello i: %v\n", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

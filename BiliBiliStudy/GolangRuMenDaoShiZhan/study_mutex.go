package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	lock.Lock()
	i += 1
	defer lock.Unlock()
	defer wg.Done()
	time.Sleep(time.Microsecond * 10)
}

func sub() {
	lock.Lock()
	defer lock.Unlock()
	i -= 1
	defer wg.Done()
	time.Sleep(time.Microsecond * 2)
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()

	fmt.Printf("end i: %v\n", i)
}

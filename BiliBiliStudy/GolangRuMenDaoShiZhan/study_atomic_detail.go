package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i: %v\n", i)

}

func test_load_store() {
	var i int32 = 100
	i2 := atomic.LoadInt32(&i)
	fmt.Printf("i2: %v\n", i2)

	atomic.StoreInt32(&i, 200)
	fmt.Printf("i: %v\n", i)
}

func test_cas() {
	var i int32 = 100
	ok := atomic.CompareAndSwapInt32(&i, 100, 200)
	if ok {
		fmt.Println("step1 ok...")
	} else {
		fmt.Println("step2 fail...")
	}
	fmt.Printf("i: %v\n", i)

	ok = atomic.CompareAndSwapInt32(&i, 100, 500)
	if ok {
		fmt.Println("step 2 ok...")
	} else {
		fmt.Println("step 2 fail...")
	}
	fmt.Printf("i: %v\n", i)
}

func main() {

}

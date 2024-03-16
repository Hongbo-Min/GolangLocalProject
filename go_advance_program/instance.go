package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	val int
}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{val: 28}
	})
	return instance
}

func main() {
	ins := Instance()
	fmt.Printf("ins.val: %v\n", ins.val)
}

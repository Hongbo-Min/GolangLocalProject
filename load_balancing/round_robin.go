package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	services := []string{"service1", "service2", "service3"}
	currentIndex := 0
	mu := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			currentIndex = (currentIndex + 1) % len(services)
			selectedService := services[currentIndex]
			fmt.Printf("seleted service : %s\n", selectedService)
		}(i)
	}
	wg.Wait()
}

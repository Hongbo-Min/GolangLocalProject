package main

import (
	"fmt"
	"sync"
	"time"
)

type Server struct {
	Name    string
	Weight  int
	current int
}

func main() {
	servers := []Server{
		{Name: "service1", Weight: 1},
		{Name: "service2", Weight: 2},
		{Name: "service3", Weight: 3},
	}
	totalWeight := 0
	for _, server := range servers {
		totalWeight += server.Weight
	}
	mu := &sync.Mutex{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			for i := range servers {
				servers[i].current += servers[i].Weight
			}
			maxIndex := 0
			maxCurrent := 0
			for i, server := range servers {
				if server.current > maxCurrent {
					maxCurrent = server.current
					maxIndex = i
				}
			}
			fmt.Printf("selected server : %s\n", servers[maxIndex].Name)
			servers[maxIndex].current -= totalWeight
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}

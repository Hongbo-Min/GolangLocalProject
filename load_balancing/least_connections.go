package main

import (
	"fmt"
	"sync"
)

type Server struct {
	Name  string
	conns int
}

func selectServer(servers *[]Server) {
	minConns := 10000
	minIndex := 0
	for i, server := range *servers {
		if server.conns < minConns {
			minIndex = i
			minConns = server.conns
		}
	}
	(*servers)[minIndex].conns++
	fmt.Printf("selected service : %s\n", (*servers)[minIndex].Name)
}

func main() {
	servers := []Server{
		{Name: "service1", conns: 0},
		{Name: "service2", conns: 0},
		{Name: "service3", conns: 0},
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			selectServer(&servers)
		}(i)
	}
	wg.Wait()
}

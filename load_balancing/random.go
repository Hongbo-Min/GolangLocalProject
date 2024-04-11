package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	services := []string{"service1", "service2", "service3"}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		selectedService := services[rand.Intn(len(services))]
		fmt.Printf("select service : %s\n", selectedService)
	}
}

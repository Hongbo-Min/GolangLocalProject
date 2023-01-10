package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	/* counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker...")
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	} */
	chInt := make(chan int)
	go func() {
		for _ = range ticker.C {
			select {
			case chInt <- 1:
			case chInt <- 2:
			case chInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chInt {
		fmt.Printf("v: %v\n", v)
		sum += v
		if sum >= 10 {
			fmt.Printf("sum: %v\n", sum)
			break
		}
	}
}

// Timer只执行一次，Ticker可以周期的执行

package main

import "fmt"

var ch = make(chan int)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	/* for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("value: %v\n", value)
	} */

	/* for {
		v, ok := <-ch
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	} */

	for v := range ch {
		fmt.Printf("v: %v\n", v)
	}
}

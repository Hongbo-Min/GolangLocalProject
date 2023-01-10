package main

import (
	"fmt"
	"time"
)

// Timer顾名思义就是定时器的意思

func main() {
	/* timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())

	t1 := <-timer.C
	fmt.Printf("t1: %v\n", t1)
	*/

	/* fmt.Printf("time.Now(): %v\n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("time.Now(): %v\n", time.Now()) */

	/* fmt.Printf("time.Now(): %v\n", time.Now())
	time.Sleep(time.Second)
	fmt.Printf("time.Now(): %v\n", time.Now()) */

	/* 	<-time.After(time.Second * 2)
	   	fmt.Println("...") */
	/* timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C
		fmt.Println("func...")
	}()
	s := timer.Stop()
	if s {
		fmt.Println("stop...") // stop之后 上面的输出func也不会执行
	}
	time.Sleep(time.Second * 3)
	fmt.Println("main...") */

	timer := time.NewTimer(time.Second * 5)
	timer.Reset(time.Second * 1)
	fmt.Println("main end...")
}

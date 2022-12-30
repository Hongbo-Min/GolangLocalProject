package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()
	var s, sep string
	for index, arg := range os.Args {
		if index == 0 {
			fmt.Println("args[0] : ", arg)
		} else {
			s += sep + arg
			sep = " "
		}
	}

	fmt.Println("s : ", s)
	fmt.Println("time sub1 : ", time.Since(timeStart).Seconds())

	timeStart = time.Now()
	fmt.Println("s : ", strings.Join(os.Args[1:], " "))
	fmt.Println("time sub2 : ", time.Since(timeStart).Seconds())

	// for index, arg := range os.Args {
	// 	fmt.Println("index: ", index, " arg: ", arg)
	// }
}

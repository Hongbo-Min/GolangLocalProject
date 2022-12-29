package main

import (
	"fmt"
	"os"
)

func main() {
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
}

package main

import "fmt"

func main() {
	a := 5
	b := 8.1
	fmt.Println(a + int(b))

	aa := [2]int{5, 6}
	bb := [3]int{5, 6}
	if aa == bb {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

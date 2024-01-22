package main

import "fmt"

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)

	fmt.Printf("add(1, 2): %v\n", add(1, 2))
	fmt.Printf("add(1, 3, 7): %v\n", add(1, 3, 7))
	// add([]int{1, 2})
	fmt.Printf("add([]int{1, 3, 7}...): %v\n", add([]int{1, 3, 7}...))
}

func incr(p *int) int {
	*p++
	return *p
}

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

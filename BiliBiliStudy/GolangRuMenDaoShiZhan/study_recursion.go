package main

import "fmt"

// 递归函数

func f1() {
	var s int = 1

	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func recursion(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursion(value-1)
	}
}

func fib(value int) int {
	if value == 1 || value == 2 {
		return value
	} else {
		return fib(value-1) + fib(value-2)
	}
}

func main() {
	fmt.Printf("recursion(5): %v\n", recursion(5))
	fmt.Printf("recursion(10): %v\n", recursion(10))

	fmt.Printf("fib(8): %v\n", fib(8))
	fmt.Printf("fib(10): %v\n", fib(10))
}

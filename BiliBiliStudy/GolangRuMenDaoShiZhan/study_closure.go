package main

import (
	"fmt"
	"strings"
)

// 闭包函数

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func makePrefixFunc(prefix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, prefix) {
			return prefix + name
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}

	sub := func(a int) int {
		base -= a
		return base
	}

	return add, sub
}

func main() {
	var f = add()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(20): %v\n", f(20))
	fmt.Printf("f(30): %v\n", f(30))
	fmt.Println("------")
	f1 := add()
	fmt.Printf("f1(10): %v\n", f1(10))
	fmt.Printf("f1(20): %v\n", f1(20))
	fmt.Printf("f1(30): %v\n", f1(30))

	checkPrefix := makePrefixFunc("http://")
	checkSuffix := makeSuffixFunc(".go")

	fmt.Printf("checkPrefix(\"www.baidu.com\"): %v\n", checkPrefix("www.baidu.com"))
	fmt.Printf("checkPrefix(\"http://www.baidu.com\"): %v\n", checkPrefix("http://www.baidu.com"))

	fmt.Printf("checkSuffix(\"main.java\"): %v\n", checkSuffix("main.java"))
	fmt.Printf("checkSuffix(\"main.go\"): %v\n", checkSuffix("main.go"))

	add, sub := calc(10)

	fmt.Println(add(5), sub(9))
	fmt.Println(add(8), sub(7))
}

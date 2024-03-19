package main

import "fmt"

func main() {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 先计算出参数，然后再将defer入栈
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

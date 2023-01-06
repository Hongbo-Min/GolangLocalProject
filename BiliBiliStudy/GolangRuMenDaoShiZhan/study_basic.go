package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func GetNameAndAge() (string, int) {
	return "tom", 20
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Cal(operactor string) func(int, int) int {
	switch operactor {
	case "+":
		return Add
	case "-":
		return Sub
	default:
		return nil
	}
}

func main() {
	value1, value2 := 10, 20
	fmt.Printf("Cal(\"+\")(value1, value2): %v\n", Cal("+")(value1, value2))
	fmt.Printf("Cal(\"-\")(value1, value2): %v\n", Cal("-")(value1, value2))

	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	// fmt.Printf("max: %v\n", max)

	fmt.Printf("max(value1, value2): %v\n", max(value1, value2))
	// const w, h = 200, 300
	// fmt.Printf("h: %v\n", h)
	// fmt.Printf("w: %v\n", w)

	// a := 100
	// p := &a
	// fmt.Printf("p: %T\n", p)

	// x := [...]int{1, 2, 3} // 数组类型
	// fmt.Printf("x: %T\n", x)

	// y := []int{1, 2, 3} // 切片类型
	// fmt.Printf("y: %T\n", y)

	a := 10

	fmt.Printf("a: %d\n", a)
	fmt.Printf("a: %b\n", a)
	fmt.Printf("a: %o\n", a)
	fmt.Printf("a: %x\n", a)

	b := 0xff
	fmt.Printf("b: %x\n", b)
	fmt.Printf("b: %X\n", b)

	fmt.Printf("math.Pi: %f\n", math.Pi)
	fmt.Printf("math.Pi: %.2f\n", math.Pi)

	name := "tom"
	age := "20"

	s := strings.Join([]string{name, age}, ",")
	fmt.Printf("s: %v\n", s)

	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("cat")
	fmt.Printf("buffer.String(): %v\n", buffer.String())
}

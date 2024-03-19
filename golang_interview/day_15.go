package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
	if s2 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}

	i := 65
	fmt.Println(string(i))
	fmt.Println(strconv.Itoa(i))

	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

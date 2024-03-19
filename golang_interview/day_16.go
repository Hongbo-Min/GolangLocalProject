package main

import "fmt"

type AA interface {
	ShowA() int
}

type BB interface {
	ShowB() int
}

type WorkW struct {
	i int
}

func (w WorkW) ShowA() int {
	return w.i + 10
}

func (w WorkW) ShowB() int {
	return w.i + 20
}

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	println(len(a), cap(a))
	println(len(b), cap(b))
	println(len(c), cap(c))

	work := WorkW{i: 3}
	var aa AA = work
	var bb BB = work
	fmt.Println(aa.ShowA())
	fmt.Println(bb.ShowB())
}

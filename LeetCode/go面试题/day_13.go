package main

import "fmt"

func main() {
	i := 5
	defer hello(i)
	i = i + 10

	p := Teacher{}
	p.ShowA()
}

func hello(i int) {
	fmt.Println(i)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

package main

import "fmt"

func test() {
	type Person struct {
		id   int
		name string
		age  int
	}
	tom := Person{
		id:   101,
		name: "tom",
		age:  20,
	}
	var p_person *Person
	p_person = &tom

	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %v\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
}

func test2() {
	type Person struct {
		id   int
		name string
		age  int
	}
	tom := new(Person)
	tom.id = 101
	tom.name = "tom"
	tom.age = 20
	fmt.Printf("tom: %v\n", tom)
}

type Person struct {
	id   int
	name string
	age  int
	Dog
}

type Dog struct {
	name  string
	color string
	age   int
}

func test3(obj *Person) {
	obj.name = "irving"
}

func main() {
	tom := Person{
		id:   1,
		name: "tom",
		age:  20,
		Dog: Dog{
			name:  "第九",
			color: "black&white",
			age:   1,
		},
	}
	fmt.Printf("tom: %v\n", tom)
}

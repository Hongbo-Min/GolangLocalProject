package main

import "fmt"

type Person struct {
	name string
}

func (p Person) eat() {
	fmt.Printf("name: %v eatting...\n", p.name)
}

func (p Person) sleep() {
	fmt.Printf("name: %v sleeping...\n", p.name)
}

type Customer struct {
	name string
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == customer.name && pwd == "123456" {
		return true
	} else {
		return false
	}
}

func (per Person) showPerson1() {
	per.name = "tom ..."
}

func (per *Person) showPerson2() {
	// (*per).name = "tom ..."

	per.name = "tom ..." // 自动解引用
}

func main() {
	tom := Person{
		name: "tom",
	}
	tom.eat()
	tom.sleep()

	customer := Customer{
		name: "tom",
	}

	fmt.Printf("customer.login(): %v\n", customer.login("tom", "123456"))

	fmt.Printf("tom: %T\n", tom)
	tom2 := &tom
	fmt.Printf("tom2: %T\n", tom2)

	p1 := Person{
		name: "tom",
	}
	p2 := &Person{
		name: "tom",
	}

	p1.showPerson1()
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("-------")
	p2.showPerson2()
	fmt.Printf("p2: %v\n", p2)
}

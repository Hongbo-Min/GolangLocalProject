package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog struct {
	Animal
	color string
}

type Keji struct {
	Dog
	brand string
}

func main() {
	keji := Keji{
		Dog: Dog{
			Animal: Animal{
				name: "小柯基",
				age:  10,
			},
			color: "白色",
		},
		brand: "短腿",
	}

	keji.eat()
	keji.sleep()
	fmt.Printf("keji: %v\n", keji)
}

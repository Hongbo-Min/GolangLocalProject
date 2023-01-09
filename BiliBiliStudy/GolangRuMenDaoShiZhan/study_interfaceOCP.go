package main

import "fmt"

// OCP原则  开-闭原则 对扩展开放，对修改关闭

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (d Dog) eat() {
	fmt.Println("dog eat...")
}

func (d Dog) sleep() {
	fmt.Println("dog sleep...")
}

type Cat struct {
	name string
	age  int
}

func (c Cat) eat() {
	fmt.Println("cat eat...")
}

func (c Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
	name string
}

func (p Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	per := Person{}

	per.care(dog)
	per.care(cat)
}

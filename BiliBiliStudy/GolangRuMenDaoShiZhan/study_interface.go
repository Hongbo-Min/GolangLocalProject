package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read...")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}

func (m Mobile) read() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("read...")
}

func (m Mobile) write() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("write...")
}

type Pet interface {
	eat(string) string
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (d Dog) eat(name string) string {
	d.name = name
	return "dog: " + d.name + "吃的噶胖"
}

func (c Cat) eat(name string) string {
	c.name = name
	return "cat: " + c.name + "吃的噶胖"
}

type Player interface {
	playMusic(string)
}

type Video interface {
	playVideo(string)
}

type Mobile struct {
	model string
}

func (m Mobile) playMusic(name string) {
	fmt.Printf("play music: %v\n", name)
}

func (m Mobile) playVideo(name string) {
	fmt.Printf("play video: %v\n", name)
}

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (fs Fish) fly() {
	fmt.Println("fly...")
}

func (fs Fish) swim() {
	fmt.Println("swim...")
}

func main() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()

	c := Computer{
		name: "mac",
	}

	c.read()
	c.write()

	m := Mobile{
		model: "iphone",
	}
	m.read()
	m.write()

	cat := Cat{
		name: "dj",
	}
	fmt.Printf("cat.eat(\"第九\"): %v\n", cat.eat("第九"))

	m.playMusic("最后的人")
	m.playVideo("乡村爱情15")

	dog := Dog{
		name: "xh",
	}

	fmt.Printf("dog.eat(\"小黑\"): %v\n", dog.eat("小黑"))

	var pet Pet
	pet = Dog{
		name: "小黑",
	}
	fmt.Printf("pet.eat(\"小黑\"): %v\n", pet.eat("小黑"))
	pet = Cat{
		name: "第九",
	}
	fmt.Printf("pet.eat(\"第九\"): %v\n", pet.eat("第九"))
}

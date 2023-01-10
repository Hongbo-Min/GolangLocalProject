package main

import (
	"fmt"
	"log"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	} else if age < 0 {
		return nil, fmt.Errorf("age 不能小于0")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	per, err := NewPerson("第九", 1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("per: %v\n", *per)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
		Name string
		Type int `json:"int" id:"100"`
	}
	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"))
		fmt.Println(catType.Tag.Get("id"))
	}
}

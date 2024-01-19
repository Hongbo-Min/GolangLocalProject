package main

import "fmt"

func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		// return
	} else {
		fmt.Println("not nil")
	}

	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("nonempty(data): %v\n", nonempty(data))
	fmt.Printf("data: %v\n", data)
}

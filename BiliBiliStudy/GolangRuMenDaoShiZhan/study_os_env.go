package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("os.Environ(): %v\n", os.Environ())
	fmt.Printf("os.Getenv(\"GOPATH\"): %v\n", os.Getenv("GOPATH"))
	fmt.Printf("os.Getenv(\"GOROOT\"): %v\n", os.Getenv("GOROOT"))

	s, b := os.LookupEnv("GOPATH")
	if b {
		fmt.Printf("s: %v\n", s)
	}

	// os.Setenv("min", "hongbo")

	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	os.Clearenv()
}

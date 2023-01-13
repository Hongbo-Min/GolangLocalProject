package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// r := strings.NewReader("hello")
	f, _ := os.Open("study_ioutil.go")
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("string(b): %v\n", string(b))
	}
}

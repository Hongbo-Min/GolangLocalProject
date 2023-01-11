package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/* r := strings.NewReader("hello")
	buf := make([]byte, 100)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("n: %v\n", n)
	} */

	/* r := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	if _, err := io.CopyBuffer(os.Stdout, r, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	} */

	pr, pw := io.Pipe()
	go func() {
		fmt.Fprintf(pw, "some io.Reader stream to be read\n")
		pw.Close()
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

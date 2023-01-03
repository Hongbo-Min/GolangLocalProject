package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "quit" {
			break
		} else {
			if !seen[line] {
				seen[line] = true
				fmt.Println(line)
			}
		}
	}

	for k, v := range seen {
		fmt.Println("key: ", k, " value: ", v)
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

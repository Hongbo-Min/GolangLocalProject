package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	sha384Flag := flag.Bool("sha384", false, "Use SHA384 hash algorithm")
	sha512Flag := flag.Bool("sha512", false, "Use SHA512 hash algorithm")
	flag.Parse()

	var hasher func([]byte) []byte
	if *sha384Flag {
		hasher = func(data []byte) []byte {
			hash := sha512.New384()
			hash.Write(data)
			return hash.Sum(nil)
		}
	} else if *sha512Flag {
		hasher = func(data []byte) []byte {
			hash := sha512.New()
			hash.Write(data)
			return hash.Sum(nil)
		}
	} else {
		hasher = func(data []byte) []byte {
			hash := sha256.Sum256(data)
			return hash[:]
		}
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		os.Exit(1)
	}

	hash := hasher(input)
	fmt.Printf("%x\n", hash)
}
